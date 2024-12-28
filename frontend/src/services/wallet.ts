import { api } from "../config/api";

interface GetWalletResponse {
    success: boolean;
    message: string;
    error: string | null;
    data: {
        token: number;
    } | null;
}

interface TopupRequest {
    amount: number;
}

interface TopupResponse {
    success: boolean;
    message: string;
    error: string | null;
    data: {
        token: number;
    } | null;
}

export const walletService = {
    async getWallet() {
        const { data } = await api.get<GetWalletResponse>('/wallet');
        return data;
    },

    async topup(req: TopupRequest) {
        const { data } = await api.post<TopupResponse>('/topup', req);
        return data;
    },
}