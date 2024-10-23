import { applyTo, comment, initClients, postUpdate } from "./contract";
import { loadOnboard } from "./onboard";
import { chainIdHex } from "./chain";

export class ApplyForm extends HTMLElement {
    form: HTMLFormElement | null = null;
    success: HTMLDivElement | null = null;
    error: HTMLSpanElement | null = null;

    connectedCallback() {
        const id = this.getAttribute('project-id');
        if (!id) throw new Error('project-id is required');

        this.form = this.querySelector('form');
        if (!this.form) throw new Error("form child required");
        this.success = this.querySelector('[data-success]');
        this.error = this.querySelector('[data-error]');

        this.form.addEventListener('submit', async (event) => {
            event.preventDefault();
            this.success.hidden = true;
            this.error.hidden = true;
            const value = this.form.querySelector('textarea')!.value;
            const { signature, address, message } = await this.signMessage(id, value);
            const saved = await this.save(id, signature, address, message);
            if (!saved) return;

            const sent = await this.sendTx(id);
            if (!sent) return;

            this.success.hidden = false;
            this.error.hidden = true;
            this.form.classList.add('hidden');
        });
    }

    async signMessage(id: string, value: string) {
        const onboard = await loadOnboard();
        const wallets = await onboard.connectWallet()
        if (wallets.length === 0) return;

        const rightChain = await onboard.setChain({ chainId: chainIdHex })
        if (!rightChain) return;

        const { client, address } = await initClients();

        const message = JSON.stringify({
            projectId: id,
            teamApplicationProposal: true,
            value
        });

        const signature = await client.signMessage({ account: address, message });
        return { signature, address, message };
    }

    async save(id: string, signature: string, address: string, value: string) {
        try {
            const response = await fetch(`/projects/${id}/apply`, {
                method: 'post',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ signature, address, value })
            });
            if (!response.ok) {
                this.error.hidden = false;
                return false;
            }
        } catch (err) {
            console.error(err);
            this.error.hidden = false;
            return false;
        }

        return true;
    }

    async sendTx(id: string) {
        const onboard = await loadOnboard();
        const wallets = await onboard.connectWallet()
        if (wallets.length === 0) return;

        const rightChain = await onboard.setChain({ chainId: chainIdHex })
        if (!rightChain) return;

        const success = await applyTo(id);
        if (!success) {
            this.error.hidden = false;
        }
        return success;
    }
}