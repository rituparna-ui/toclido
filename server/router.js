const express = require("express");
const router = express.Router();

const authMiddleware = require("./middleware/auth");

const authRoutes = require("./routers/auth");
const todoRoutes = require("./routers/todos");
const sshRoutes = require("./routers/ssh-keys");
const cliRoutes = require("./routers/cli");
const { getUserTokenFromKey } = require("./controllers/ssh-keys");

router.use("/auth", authRoutes);

router.use("/todos", authMiddleware(), todoRoutes);

router.post("/ssh-keys/token", getUserTokenFromKey);

router.use("/ssh-keys", authMiddleware(), sshRoutes);

router.use("/cli-tokens", authMiddleware(), cliRoutes);

module.exports = router;
