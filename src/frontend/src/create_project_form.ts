import { createProject } from "./contract";
import { loadOnboard } from "./onboard";
import { chainIdHex } from "./chain";

export class CreateProjectForm extends HTMLElement {
    errorAlert: HTMLDivElement | null = null;

    cid = '';
    shortName = '';

    connectedCallback() {
        console.log("Network", process.env.NETWORK);

        this.errorAlert = this.querySelector('#error')
        const form = this.querySelector('form');
        const nextStep = this.querySelector<HTMLDivElement>('[data-step="2"]');

        if (!form) throw new Error("form child required");
        if (!nextStep) throw new Error("step 2 child required");

        form.addEventListener('submit', async (event) => {
            event.preventDefault();
            this.hideError();

            const response = await fetch(form.action, { method: form.method, body: new URLSearchParams(new FormData(form) as any) });
            if (!response.ok) {
                // try {
                //     const body = await response.json();
                //     if (body['validation']) {
                //         this.showValidation(body['validation']);
                //     }
                // } catch (error) {
                //     this.showError('We are having issues processing the form, please try again in a minute');
                //     return;
                // }
                this.showError('We are having issues processing the form, please try again in a minute');
                return;
            }

            const body = await response.json();

            this.cid = body['cid'] as string;
            if (!this.cid) throw new Error("no cid");

            this.shortName = form.querySelector<HTMLInputElement>('[name="shortName"]')!.value;

            form.classList.add('hidden');
            nextStep.hidden = false;
        });

        nextStep.querySelector('button:not(.primary)')!.addEventListener('click', () => {
            form.classList.remove('hidden');
            nextStep.hidden = true;
        });
        nextStep.querySelector('button.primary')!.addEventListener('click', () => this.sendTx());
    }

    showError(error: string) {
        if (!this.errorAlert) {
            console.error(error);
            return;
        }

        this.errorAlert.textContent = error;
        this.errorAlert.removeAttribute('hidden');
    }

    hideError() {
        if (!this.errorAlert) return;
        this.errorAlert.hidden = true;
    }

    async sendTx() {
        const onboard = await loadOnboard();
        const wallets = await onboard.connectWallet()
        if (wallets.length === 0) return;

        const rightChain = await onboard.setChain({ chainId: chainIdHex })
        if (!rightChain) return;

        const [success, projectId] = await createProject(this.cid, this.shortName);
        if (!success) {
            this.showError('Something went wrong');
        }
        window.location.pathname = `/projects/${projectId.toString()}`
    }

    // showValidation(messages: Record<string, string>) {
    //     const fields = this.querySelectorAll('label');
    //     fields.forEach((field) => {
    //         const input = field.querySelector<HTMLInputElement|HTMLTextAreaElement>('input, textarea');
    //         if (!input) return;

    //         const message = messages[input.name];
    //         if (!message) return;

    //         const element = field.querySelector('.validation')
    //         if (!element) return;

    //         element.textContent = message;
    //     });
    // }
}
