import axios from 'axios';
import { CreateServer, LoginData, RegisterData, UpdateServer } from '../models';
import Cookies from 'js-cookie';

const API_URL = process.env.NEXT_PUBLIC_API_URL;

async function register(formData: RegisterData) {
  const jsonData = {
    email: formData.email,
    username: formData.username,
    password: formData.password,
    confirm_password: formData.confirmPassword,
  };

  const response = await axios.post(`${API_URL}/signup`, jsonData);

  setCurrentUser(response);

  return response.data;
}

async function login(formData: LoginData) {
  const response = await axios.post(`${API_URL}/login`, formData, {
    withCredentials: true,
  });

  setCurrentUser(response);

  return response.data;
}

async function logout() {
  const response = await axios.get(`${API_URL}/logout`, {
    withCredentials: true,
  });

  Cookies.remove('user');
  return response.data;
}

function setCurrentUser(axiosResponse: any) {
  const { id, username, email } = axiosResponse.data;

  if (axiosResponse.data) {
    Cookies.set('user', JSON.stringify({ id, username, email }));
  }
}

async function getUserById(id: string) {
  const response = await axios.get(`${API_URL}/users/${id}`, {
    withCredentials: true,
  });
  return response.data;
}

async function updateEmailNotifications() {
  const response = await axios.patch(`${API_URL}/notification-preferences`, null, {
    withCredentials: true,
  });
  return response.data;
}

async function getServers() {
  const response = await axios.get(`${API_URL}/remote-servers`, {
    withCredentials: true,
  })
  return response.data;
}

async function createServer(formData: CreateServer) {
  const cookie = Cookies.get("user");
  const { id } = JSON.parse(cookie || "{}");
  formData.user_id = Number(id);
  const response = await axios.post(`${API_URL}/remote-servers`, formData, {
    withCredentials: true,
  })
  return response.data;
}

async function updateServer(formData: UpdateServer, id: string) {
  const response = await axios.put(`${API_URL}/remote-servers/${id}`, formData, {
    withCredentials: true,
  })
  return response.data;
}

async function getServerById(id: string) {
  const response = await axios.get(`${API_URL}/remote-servers/${id}`, {
    withCredentials: true,
  })
  return response.data;
}

async function startMonitoring(id: string) {
  const response = await axios.post(`${API_URL}/remote-servers/${id}/start-monitoring`, null, {
    withCredentials: true,
  })
  return response.data;
}

export { register, login, logout, getUserById, updateEmailNotifications, getServers, createServer, getServerById, startMonitoring, updateServer};
