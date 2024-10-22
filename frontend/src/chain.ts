import { arbitrum, foundry, arbitrumSepolia } from 'viem/chains'

export const viemChain = (() => {
    if (process.env.NETWORK === "test") {
        return foundry;
    }

    if (process.env.NETWORK === "arb-sepolia") {
        return arbitrumSepolia;
    }

    return arbitrum;
})();

export const onboardChain = [{
    id: viemChain.id,
    token: viemChain.nativeCurrency.symbol,
    label: viemChain.name,
    rpcUrl: viemChain.rpcUrls.default.http[0],
}];

export const chainIdHex = `0x${viemChain.id.toString(16)}`;
