const express = require("express");
const {
  createSshKey,
  getSshKeys,
  getUserTokenFromKey,
} = require("../controllers/ssh-keys");
const router = express.Router();

router.post("/", createSshKey);

router.get("/", getSshKeys);

router.post("/token", getUserTokenFromKey);

module.exports = router;
