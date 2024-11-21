export default function useTodos() {
  const todos = useState("todos", () => []);

  return { todos };
}
