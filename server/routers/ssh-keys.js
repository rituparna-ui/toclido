const express = require("express");
const { createSshKey, getSshKeys } = require("../controllers/ssh-keys");
const router = express.Router();

router.post("/", createSshKey);

router.get("/", getSshKeys);

module.exports = router;
