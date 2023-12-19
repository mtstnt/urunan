import { create } from "zustand";

type FlashState = {
    data: Map<string, string>
    addFlashMessage(key: string, value: string): void
    hasFlashMessage(key: string): boolean
    readFlashMessage(key: string): string
};

export const useFlashMessages = create<FlashState>((set, get) => ({
    data: new Map<string, string>(),
    addFlashMessage: (key: string, value: string) => {
        set((state) => ({ data: { ...state.data, [key]: value } }));
    },
    hasFlashMessage: (key: string): boolean => get().data.has(key),
    readFlashMessage: (key) => {
        const currentState = get();
        if (currentState.data.has(key)) {
            throw new Error();
        }
        const value = currentState.data.get(key)!;
        currentState.data.delete(key);
        return value;
    }
}));