const express = require("express");
const {
  createSshKey,
  getSshKeys,
  getUserTokenFromKey,
  deleteSshKey,
} = require("../controllers/ssh-keys");
const router = express.Router();

router.post("/", createSshKey);

router.get("/", getSshKeys);

router.post("/token", getUserTokenFromKey);

router.delete("/:id", deleteSshKey);

module.exports = router;
