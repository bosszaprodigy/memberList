import axios from 'axios'

const API_BASE = 'http://localhost:3000/api'

async function getAll() {
  const res = await axios.get(`${API_BASE}/members`)
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
