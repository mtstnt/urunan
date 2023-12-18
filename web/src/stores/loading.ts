import { create } from "zustand";

type LoadingStore = {
    isLoading: boolean,
    setIsLoading: (b: boolean) => void
};

export const useLoadingStore = create<LoadingStore>((set) => ({
    isLoading: false,
    setIsLoading: (b: boolean) => set(() => ({ isLoading: b })),
}));