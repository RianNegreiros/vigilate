import axios from 'axios';
import { LoginData, RegisterData } from '../models';

const API_URL = process.env.NEXT_PUBLIC_API_URL;

async function register(formData: RegisterData) {
  const jsonData = {
    email: formData.email,
    username: formData.username,
    password: formData.password,
    confirm_password: formData.confirmPassword,
  };

  const response = await axios.post(`${API_URL}/signup`, jsonData);
  return response.data;
}

async function login(formData: LoginData) {
  const response = await axios.post(`${API_URL}/login`, formData);
  return response.data;
}

async function logout() {
  const response = await axios.get(`${API_URL}/logout`);
  return response.data;
}

export { register, login, logout };
