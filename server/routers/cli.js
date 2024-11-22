const express = require("express");

const {
  getAllTokens,
  createCliToken,
  deleteCliToken,
} = require("../controllers/cli");

const router = express.Router();

router.get("/", getAllTokens);

router.post("/", createCliToken);

router.post("/revoke", deleteCliToken);

module.exports = router;
