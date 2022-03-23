import TableRow from "@mui/material/TableRow"
import TableCell from "@mui/material/TableCell"
import ButtonGroup from "@mui/material/ButtonGroup"
import Button from "@mui/material/Button"

const Todo = ({ id, title, finished, onFinished, onDelete }) => {
	return (
		<TableRow key={id}>
			<TableCell>{id}</TableCell>
			<TableCell>{title}</TableCell>
			<TableCell>{finished ? "Finished" : "Not Yet"}</TableCell>
			<TableCell>
				<ButtonGroup>
					<Button disabled={finished} color="success" onClick={onFinished}>
						Finish
					</Button>
					<Button onClick={onDelete} color="error">
						Delete
					</Button>
				</ButtonGroup>
			</TableCell>
		</TableRow>
	)
}

export default Todo
