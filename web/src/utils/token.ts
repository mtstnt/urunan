
const STORAGE_KEY = "token";

export function getTokenFromStorage(): string | null {
    return localStorage.getItem(STORAGE_KEY);
}

export function storeTokenInStorage(token: string) {
    localStorage.setItem(STORAGE_KEY, token);
}

export async function wait(n: number) {
    return new Promise<void>((resolve) => {
        const t = setTimeout(() => {
            clearTimeout(t);
            resolve();
        }, n * 1000);
    });
}