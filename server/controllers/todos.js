const Todo = require("./../models/todo");

exports.getTodos = async (req, res) => {
  try {
    const todos = (
      await Todo.find({ user: req.user, isDeleted: false })
        .sort({ _id: -1 })
        .lean()
    ).map((todo, i) => ({
      ...todo,
      i: i + 1,
    }));
    return res.json({
      todos,
    });
  } catch (error) {
    return res.status(500).json({
      message: "Internal Server",
      errorMessage: error.message,
      error,
    });
  }
};

exports.createTodo = async (req, res) => {
  try {
    const { title, description } = req.body;
    if (!title || !description) {
      return res.status(400).json({ message: "Please provide all the fields" });
    }
    await Todo.create({
      title,
      description,
      user: req.user,
    });
    return res.json({
      message: "Todo created",
    });
  } catch (error) {
    return res.status(500).json({
      message: "Internal Server",
      errorMessage: error.message,
      error,
    });
  }
};
