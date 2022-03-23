import { useEffect, useState } from "react"
import Table from "@mui/material/Table"
import TableBody from "@mui/material/TableBody"
import TableCell from "@mui/material/TableCell"
import TableContainer from "@mui/material/TableContainer"
import TableHead from "@mui/material/TableHead"
import TableRow from "@mui/material/TableRow"
import TextField from "@mui/material/TextField"
import Typography from "@mui/material/Typography"
import Container from "@mui/material/Container"
import Box from "@mui/material/Box"
import * as api from "./api"
import Todo from "./components/Todo"
import { Button } from "@mui/material"

function App() {
	const [todos, setTodos] = useState([])
	const [title, setTitle] = useState("")
	const [isError, setIsError] = useState(false)

	useEffect(() => {
		fetchTodos()
	}, [])

	const fetchTodos = () => {
		api
			.getTodos()
			.then((fetchedTodos) => setTodos(fetchedTodos))
			.catch((err) => console.log(err))
	}

	const onCreateTodo = () => {
		if (title.length === 0) {
			setIsError(true)
			return
		}
		api
			.createTodo(title)
			.then(() => {
				fetchTodos()
				setIsError(false)
				setTitle("")
			})
			.catch((err) => console.log(err))
	}

	const finishTodo = (id) => {
		api
			.finishTodo(id)
			.then(() => fetchTodos())
			.catch((err) => console.log(err))
	}

	const deleteTodo = (id) => {
		api
			.deleteTodo(id)
			.then(() => fetchTodos())
			.catch((err) => console.log(err))
	}

	return (
		<Container>
			<Typography variant="h1">Todo List</Typography>
			<Box sx={{ m: 8 }} />
			<div>
				<TextField
					error={isError}
					value={title}
					label="Todo Title"
					variant="standard"
					onChange={(e) => {
						setTitle(e.target.value)
					}}
				/>
				<Button variant="contained" onClick={onCreateTodo}>
					Create
				</Button>
			</div>
			<Box sx={{ m: 8 }} />
			<TableContainer>
				<Table>
					<TableHead>
						<TableRow>
							<TableCell>ID</TableCell>
							<TableCell>Title</TableCell>
							<TableCell>Status</TableCell>
							<TableCell>Operation</TableCell>
						</TableRow>
					</TableHead>
					<TableBody>
						{todos.map(({ ID, name, finished }) => {
							return (
								<Todo
									key={ID}
									id={ID}
									title={name}
									finished={finished}
									onDelete={() => deleteTodo(ID)}
									onFinished={() => finishTodo(ID)}
								/>
							)
						})}
					</TableBody>
				</Table>
			</TableContainer>
		</Container>
	)
}

export default App
