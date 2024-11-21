const express = require("express");
const {
  getTodos,
  createTodo,
  deleteTodo,
  toggleStatus,
} = require("../controllers/todos");
const router = express.Router();

router.get("/", getTodos);

router.post("/", createTodo);

router.delete("/:id", deleteTodo);

router.patch("/toggle-status/:id", toggleStatus);

module.exports = router;
