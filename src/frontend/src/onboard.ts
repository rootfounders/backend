import Onboard, { OnboardAPI } from '@web3-onboard/core'
import injectedModule from '@web3-onboard/injected-wallets'
import { onboardChain } from './chain';

let onboard: OnboardAPI|null = null;

export async function loadOnboard(): Promise<OnboardAPI> {
    if (onboard) {
        return onboard;
    }

    await Promise.all([
        import('@web3-onboard/core'),
        import('@web3-onboard/injected-wallets'),
    ]);

    const injected = injectedModule()

    onboard = Onboard({
        wallets: [injected],
        chains: onboardChain,
    })

    return onboard;
}