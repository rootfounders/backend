import { loadOnboard } from "./onboard";

export class CreateProject extends HTMLElement {
    connectedCallback() {
        this.innerHTML = `<button type="button" class="primary">Create Project</button>`;
        this.addEventListener('click', async (event) => {
            if (!(event.target instanceof HTMLButtonElement)) return;

            const wallets = await (await loadOnboard()).connectWallet()
            if (wallets.length > 0) {
                alert("connected");
            } else {
                alert("Nope");
            }
        });
    }
}
