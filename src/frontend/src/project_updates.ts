import { comment, initClients, postUpdate } from "./contract";
import { loadOnboard } from "./onboard";
import { chainIdHex } from "./chain";

type Update = {
    content: string
    date: number
};

export class ProjectUpdates extends HTMLElement {
    ownerAddress: string = "";
    form: HTMLFormElement | null = null;
    success: HTMLSpanElement | null = null;
    error: HTMLSpanElement | null = null;

    updateTemplate: HTMLTemplateElement | null = null;
    updates: Update[] = [];

    connectedCallback() {
        const id = this.getAttribute('project-id');
        if (!id) throw new Error('project-id is required');
        this.ownerAddress = this.getAttribute('owner');

        this.form = this.querySelector('form');
        if (!this.form) throw new Error("form child required");
        this.success = this.querySelector('[data-success]');
        this.error = this.querySelector('[data-error]');

        this.updateTemplate = this.querySelector('template#update');

        const raw = this.getAttribute("updates");
        if (!raw) return;

        this.updates = JSON.parse(raw);

        this.form.addEventListener('submit', async (event) => {
            event.preventDefault();
            this.success.hidden = true;
            this.error.hidden = true;
            const value = this.form.querySelector('textarea')!.value;
            this.sendTx(id, value);
        });

        this.querySelector('button[data-show-form]').addEventListener('click', ({ target }) => {
            this.showForm();
            // @ts-ignore
            target.disabled = true;
        });

        this.showUpdates();
    }

    async showForm() {
        const onboard = await loadOnboard();
        const wallets = await onboard.connectWallet()
        if (wallets.length === 0) return;

        await onboard.setChain({ chainId: chainIdHex });
        const { address } = await initClients();
        if (address != this.ownerAddress) {
            console.error('Only project owner can post updates');
            return;
        }

        this.form.classList.remove('hidden');
    }

    async sendTx(id: string, value: string) {
        const onboard = await loadOnboard();
        const wallets = await onboard.connectWallet()
        if (wallets.length === 0) return;

        const rightChain = await onboard.setChain({ chainId: chainIdHex })
        if (!rightChain) return;

        const success = await postUpdate(id, value);
        if (!success) {
            this.error.hidden = false;
            this.success.hidden = true;
            return
        }
        this.success.hidden = false;

        this.updates.unshift({
            content: value,
            date: Date.now(),
        });
        this.showUpdates();
    }

    showUpdates() {
        const fragment = document.createDocumentFragment();

        this.updates.forEach((update) => {
            const template = this.updateTemplate.content.cloneNode(true) as HTMLLIElement;
            template.querySelector('[data-content]').textContent = update.content;
            template.querySelector('[data-posted-date]').textContent = (new Date(update.date)).toDateString();
            fragment.appendChild(template);
        });

        this.querySelector('#updates').replaceChildren(fragment);
    }
}