const express = require("express");
const router = express.Router();

const authMiddleware = require("./middleware/auth");

const authRoutes = require("./routers/auth");
const todoRoutes = require("./routers/todos");

router.use("/auth", authRoutes);

router.use("/todos", authMiddleware(), todoRoutes);

module.exports = router;
