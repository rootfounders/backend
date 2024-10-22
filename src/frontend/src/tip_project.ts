import { tip } from "./contract";
import { loadOnboard } from "./onboard";
import { chainIdHex } from "./chain";

export class TipProject extends HTMLElement {
    success: HTMLSpanElement | null = null;

    connectedCallback() {
        const id = this.getAttribute('project-id');
        if (!id) throw new Error('project-id is required');

        const form = this.querySelector('form');
        if (!form) throw new Error("form child required");
        this.success = this.querySelector('[data-success]');

        form.addEventListener('submit', async (event) => {
            event.preventDefault();
            this.success.hidden = true;
            const value = form.querySelector('input')!.value;
            this.sendTx(id, value);
        });
    }

    async sendTx(id: string, value: string) {
        const onboard = await loadOnboard();
        const wallets = await onboard.connectWallet()
        if (wallets.length === 0) return;

        const rightChain = await onboard.setChain({ chainId: chainIdHex })
        if (!rightChain) return;

        const success = await tip(id, value);
        if (!success) {
            console.error('Something went wrong');
            return
        }
        this.success.hidden = false;
    }
}