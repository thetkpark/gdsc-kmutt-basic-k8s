import axios from "axios"

const instance = axios.create({
	baseURL: process.env.REACT_APP_API_BASEURL
})

export const getTodos = async () => {
	const res = await instance.get("/todos")
	return res.data
}

export const createTodo = async (title) => {
	const res = await instance.post("/todo", { name: title, finished: false })
	return res.data
}

export const finishTodo = async (id) => {
	const res = await instance.patch("/todo/" + id)
	return res.data
}

export const deleteTodo = async (id) => {
	const res = await instance.delete("/todo/" + id)
	return res.data
}
