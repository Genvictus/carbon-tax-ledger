import { api } from "../config/api";

interface GetCarbonResponse {
    success: boolean;
    message: string;
    error: string | null;
    data: {
        token: number;
    } | null;
}

interface PayCarbonTaxRequest {
    amount: number;
}

interface PayCarbonTaxResponse {
    success: boolean;
    message: string;
    error: string | null;
    data: {
        token: number;
    } | null;
}

export const carbonService = {
    async getCarbon() {
        const { data } = await api.get<GetCarbonResponse>('/carbon');
        return data;
    },

    async payCarbonTax(req: PayCarbonTaxRequest) {
        const { data } = await api.post<PayCarbonTaxResponse>('/pay', req);
        return data;
    },

    async getHistory() {
        const { data } = await api.get<number[]>('/history');
        return data;
    }
}