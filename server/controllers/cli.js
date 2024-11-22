const cuid = require("@paralleldrive/cuid2");

const User = require("./../models/user");
const CliToken = require("./../models/cli-token");

exports.getAllTokens = async (req, res) => {
  try {
    const tokens = await CliToken.find({
      user: req.user,
      isDeleted: false,
    }).sort({ _id: -1 });
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
};

exports.createCliToken = async (req, res) => {
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
};

exports.deleteCliToken = async (req, res) => {
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
};
