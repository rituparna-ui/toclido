const express = require("express");
const router = express.Router();

const authMiddleware = require("./middleware/auth");

const authRoutes = require("./routers/auth");
const todoRoutes = require("./routers/todos");
const sshRoutes = require("./routers/ssh-keys");
const cliRoutes = require("./routers/cli");

router.use("/auth", authRoutes);

router.use("/todos", authMiddleware(), todoRoutes);

router.use("/ssh-keys", authMiddleware(), sshRoutes);

router.use("/cli-tokens", authMiddleware(), cliRoutes);

module.exports = router;
