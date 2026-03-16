import axios from 'axios'

const API_BASE = 'http://localhost:3000/api'

async function getAll({ page, limit }) {
  const params = new URLSearchParams();

  if (page) params.append('page', page);
  if (limit) params.append('limit', limit);

  const res = await axios.get(`${API_BASE}/members?${params.toString()}`)
  return { success: true, response: res.data.data }
}

async function create(data) {
  try {
    const res = await axios.post(`${API_BASE}/members`, data)
    return { success: true, response: res.data }
  } catch (error) {
    error.value = error.response?.data?.error || ''
  }
}

export default {
  getAll,
  create,
}
