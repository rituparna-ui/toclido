const express = require("express");
const {
  createSshKey,
  getSshKeys,
  deleteSshKey,
} = require("../controllers/ssh-keys");
const router = express.Router();

router.post("/", createSshKey);

router.get("/", getSshKeys);

router.delete("/:id", deleteSshKey);

module.exports = router;
