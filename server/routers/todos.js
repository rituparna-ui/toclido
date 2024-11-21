const express = require("express");
const { getTodos, createTodo, deleteTodo } = require("../controllers/todos");
const router = express.Router();

router.get("/", getTodos);

router.post("/", createTodo);

router.delete("/:id", deleteTodo);

module.exports = router;
