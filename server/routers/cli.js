const express = require("express");
const cuid = require("@paralleldrive/cuid2");

const User = require("./../models/user");
const CliToken = require("./../models/cli-token");

const router = express.Router();

router.get("/", async (req, res) => {
  try {
    const tokens = await CliToken.find({ user: req.user, isDeleted: false });
    return res.json({
      tokens,
    });
  } catch (error) {
    return res.status(500).json({
      message: "Internal Server",
      errorMessage: error.message,
      error,
    });
  }
});

router.post("/", async (req, res) => {
  try {
    const token = cuid.createId();
    await CliToken.create({
      token,
      user: req.user,
    });
    const user = await User.findOne({ _id: req.user }).select("email");
    return res.json({
      message: "CLI Token Added",
      token,
      user: user.email,
    });
  } catch (error) {
    return res.status(500).json({
      message: "Internal Server",
      errorMessage: error.message,
      error,
    });
  }
});

router.post("/revoke", async (req, res) => {
  try {
    const { token } = req.body;
    await CliToken.findOneAndUpdate({ token }, { $set: { isDeleted: true } });
    return res.json({
      message: "CLI Token Revoked",
    });
  } catch (error) {
    return res.status(500).json({
      message: "Internal Server",
      errorMessage: error.message,
      error,
    });
  }
});

module.exports = router;
