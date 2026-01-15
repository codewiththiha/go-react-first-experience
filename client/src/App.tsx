// src/App.tsx
import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/navbar";
import TodoForm from "./components/todo-form";
import TodoList from "./components/todo-list";

export default function App() {
	return (
		<>
			<Stack>
				<Navbar />

				<Container w={700} marginTop={10}>
					<TodoForm />
					<TodoList />
				</Container>
			</Stack>
		</>
	);
}
