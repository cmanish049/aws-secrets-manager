import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
});

export interface Secret {
  name: string;
  value?: string;
  description?: string;
}

export const listSecrets = async (): Promise<Secret[]> => {
  const response = await api.get<Secret[]>('/secrets');
  return response.data || [];
};

export const getSecret = async (name: string): Promise<Secret> => {
  const response = await api.get<Secret>(`/secrets/${name}`);
  return response.data;
};

export const createSecret = async (secret: Secret): Promise<void> => {
  await api.post('/secrets', secret);
};

export const updateSecret = async (name: string, value: string): Promise<void> => {
  await api.put(`/secrets/${name}`, { value });
};

export default api;
