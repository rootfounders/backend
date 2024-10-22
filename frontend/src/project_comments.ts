import { comment, initClients } from "./contract";
import { loadOnboard } from "./onboard";
import { chainIdHex } from "./chain";

type Comment = {
    author: string
    content: string
}

export class ProjectComments extends HTMLElement {
    success: HTMLSpanElement | null = null;
    commentTemplate: HTMLTemplateElement | null = null;
    comments: Comment[] = [];

    connectedCallback() {
        const id = this.getAttribute('project-id');
        if (!id) throw new Error('project-id is required');

        const form = this.querySelector('form');
        if (!form) throw new Error("form child required");
        this.success = this.querySelector('[data-success]');

        this.commentTemplate = this.querySelector('template');

        const raw = this.getAttribute("comments");
        if (!raw) return;

        this.comments = JSON.parse(raw);

        form.addEventListener('submit', async (event) => {
            event.preventDefault();
            this.success.hidden = true;
            const value = form.querySelector('textarea')!.value;
            this.sendTx(id, value);
        });
        this.showComments();
    }

    async sendTx(id: string, value: string) {
        const onboard = await loadOnboard();
        const wallets = await onboard.connectWallet()
        if (wallets.length === 0) return;

        const rightChain = await onboard.setChain({ chainId: chainIdHex })
        if (!rightChain) return;

        const success = await comment(id, value);
        if (!success) {
            console.error('Something went wrong');
            return
        }
        this.success.hidden = false;

        const { address } = await initClients()
        this.comments.unshift({
            author: address,
            content: value,
        });
        this.showComments();
    }

    showComments() {
        const fragment = document.createDocumentFragment();

        this.comments.forEach((comment) => {
            const template = this.commentTemplate.content.cloneNode(true) as HTMLLIElement;
            template.querySelector('[data-author]').textContent = comment.author;
            template.querySelector('[data-content]').textContent = comment.content;
            fragment.appendChild(template);
        });

        this.querySelector('#comments').replaceChildren(fragment);
    }
}