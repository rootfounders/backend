import { createWalletClient, createPublicClient, http, custom, createTestClient, walletActions, publicActions, WalletClient, parseEventLogs, PublicClient, parseEther } from 'viem'
import { viemChain } from './chain'
import abi from './RootFounders.abi.json';

const mainContractAddress = process.env.MAIN_ADDRESS as `0x{string}`

export async function initClients() {
    let client: WalletClient;
    let publicClient;
    let address;

    if (process.env.NETWORK === "test") {
        client = createTestClient({
            chain: viemChain,
            mode: 'anvil',
            transport: http(),
        }).extend(walletActions);
        publicClient = client.extend(publicActions);
    } else {
        client = createWalletClient({
            chain: viemChain,
            // @ts-ignore
            transport: custom(window['ethereum'])
        });
        publicClient = createPublicClient({
            chain: viemChain,
            transport: http()
        });
    }

    [address] = await client.getAddresses()
    return { client, publicClient, address };
}

export async function createProject(cid: string, shortName: string) {
    const { client, publicClient, address } = await initClients();
    const { request } = await publicClient.simulateContract({
        address: mainContractAddress,
        account: address,
        abi,
        functionName: 'createProject',
        args: [0, cid, shortName],
    });
    const hash = await client.writeContract(request);
    const receipt = await publicClient.waitForTransactionReceipt({ hash })
    if (receipt.status !== 'success') return [false, 0n] as readonly [boolean, BigInt];

    const [log] = parseEventLogs({
        abi,
        eventName: 'ProjectCreated',
        logs: receipt.logs,
    })

    // @ts-ignore
    return [Boolean(log), log.args.id] as readonly [boolean, BigInt];
}

export async function tip(projectId: string, value: string) {
    const { client, publicClient, address } = await initClients();

    const tipJarAddress = await publicClient.readContract({
        address: mainContractAddress,
        abi,
        functionName: 'getProjectTipJar',
        args: [BigInt(projectId)],
    }) as `0x{string}`;

    // @ts-ignore
    const hash = await client.sendTransaction({
        account: address,
        to: tipJarAddress,
        value: parseEther(value),
    });
    const receipt = await publicClient.waitForTransactionReceipt({ hash })

    return receipt.status === 'success';
}

export async function comment(projectId: string, value: string) {
    const { client, publicClient, address } = await initClients();
    const { request } = await publicClient.simulateContract({
        address: mainContractAddress,
        account: address,
        abi,
        functionName: 'comment',
        args: [BigInt(projectId), value],
    });
    const hash = await client.writeContract(request);
    const receipt = await publicClient.waitForTransactionReceipt({ hash })
    if (receipt.status !== 'success') return false;

    const [log] = parseEventLogs({
        abi,
        eventName: 'Commented',
        logs: receipt.logs,
    })

    return Boolean(log);
}

export async function applyTo(projectId: string) {
    const { client, publicClient, address } = await initClients();
    const { request } = await publicClient.simulateContract({
        address: mainContractAddress,
        account: address,
        abi,
        functionName: 'applyTo',
        args: [BigInt(projectId)],
    });
    const hash = await client.writeContract(request);
    const receipt = await publicClient.waitForTransactionReceipt({ hash })
    if (receipt.status !== 'success') return false;

    const [log] = parseEventLogs({
        abi,
        eventName: 'Applied',
        logs: receipt.logs,
    })

    return Boolean(log);
}