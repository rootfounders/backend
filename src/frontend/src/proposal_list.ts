import { applyTo, comment, initClients, postUpdate } from "./contract";
import { loadOnboard } from "./onboard";
import { chainIdHex } from "./chain";

type Proposal = {
    address: string
    message: string
    created: number
}

export class ProposalList extends HTMLElement {
    error: HTMLSpanElement | null = null;
    proposalTemplate: HTMLTemplateElement | null = null;

    connectedCallback() {
        const id = this.getAttribute('project-id');
        if (!id) throw new Error('project-id is required');

        this.proposalTemplate = this.querySelector('template');
        this.error = this.querySelector('[data-error]');

        this.querySelector('button[data-connect]').addEventListener('click', async (event) => {
            this.error.hidden = true;

            const { signature } = await this.signMessage()
            const proposals = await this.load(id, signature);
            if (!proposals) return;

            this.showProposals(proposals);
        });
    }

    showProposals(proposals: Proposal[]) {
        const fragment = document.createDocumentFragment();

        proposals.forEach((proposal) => {
            const template = this.proposalTemplate.content.cloneNode(true) as HTMLDivElement;
            template.querySelector('[data-from]').textContent = proposal.address;
            template.querySelector('[data-when]').textContent = (new Date(proposal.created)).toDateString();
            template.querySelector('[data-content]').textContent = proposal.message;
            // template.querySelector('[data-accept]').addEventListener('click', () => {

            // });
            fragment.appendChild(template);
        });

        this.querySelector('#proposals').replaceChildren(fragment);
    }

    async signMessage() {
        const onboard = await loadOnboard();
        const wallets = await onboard.connectWallet()
        if (wallets.length === 0) return;

        const rightChain = await onboard.setChain({ chainId: chainIdHex })
        if (!rightChain) return;

        const { client, address } = await initClients();

        const message = "connect wallet";

        const signature = await client.signMessage({ account: address, message });
        return { signature, address, message };
    }

    async load(id: string, signature: string) {
        try {
            const url = new URL(`/projects/${id}/proposals/list`, window.location.origin);
            const searchParams = new URLSearchParams({ signature });
            url.search = searchParams.toString();
            const response = await fetch(url);
            if (!response.ok) {
                this.error.hidden = false;
                return false;
            }

            const result = await response.json() as Promise<Proposal[]>;
            return result;
        } catch (err) {
            console.error(err);
            this.error.hidden = false;
            return null;
        }
    }
}