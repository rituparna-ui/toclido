const sshpk = require("sshpk");
const jwt = require("jsonwebtoken");

const SSHKey = require("./../models/ssh-key");

exports.createSshKey = async (req, res) => {
  try {
    const { name, publicKey: publicKeyString } = req.body;
    if (!name || !publicKeyString) {
      return res.status(400).json({ message: "Please provide all the fields" });
    }
    const publicKey = publicKeyString.trim();

    let key;
    try {
      key = sshpk.parseKey(publicKey, "ssh");
    } catch (error) {
      return res.status(400).json({ message: "Invalid public key" });
    }

    const existingKey = await SSHKey.findOne({ publicKey, isDeleted: false });
    if (existingKey) {
      return res.status(400).json({ message: "Key already exists" });
    }

    const newkey = await SSHKey.create({
      name,
      publicKey,
      user: req.user,
      fingerprint: key
        .fingerprint("sha256")
        .toString("hex")
        .replaceAll(":", ""),
    });

    const parsed = {
      type: key.type,
      comment: key.comment,
      fingerprint: key.fingerprint("sha256").toString("base64"),
      size: key.size,
    };

    return res.json({
      message: "SSH Key Added",
      sshKey: {
        ...newkey.toJSON(),
        parsed,
      },
    });
  } catch (error) {
    return res.status(500).json({
      message: "Internal Server",
      errorMessage: error.message,
      error,
    });
  }
};

exports.getSshKeys = async (req, res) => {
  try {
    const sshKeys = (
      await SSHKey.find({ user: req.user, isDeleted: false })
        .sort({ _id: -1 })
        .lean()
    ).map((sshKey, i) => {
      const parsedKey = sshpk.parseKey(sshKey.publicKey, "ssh");
      const parsed = {
        type: parsedKey.type,
        comment: parsedKey.comment,
        fingerprint: parsedKey.fingerprint("sha256").toString("base64"),
        size: parsedKey.size,
      };
      return {
        ...sshKey,
        parsed,
        i: i + 1,
      };
    });
    return res.json({
      sshKeys,
    });
  } catch (error) {
    return res.status(500).json({
      message: "Internal Server",
      errorMessage: error.message,
      error,
    });
  }
};

exports.getUserTokenFromKey = async (req, res) => {
  try {
    const { fingerprint } = req.body;
    if (!fingerprint) {
      return res.status(400).json({ message: "Please provide all the fields" });
    }

    const sshKey = await SSHKey.findOne({
      fingerprint,
      isDeleted: false,
    }).populate("user", "_id name email");
    if (!sshKey) {
      return res.status(400).json({ message: "Invalid public key" });
    }

    const token = jwt.sign({ id: sshKey.user._id }, "SeCrEtKeY");
    return res.json({
      token,
      name: sshKey.user.name,
      email: sshKey.user.email,
    });
  } catch (error) {
    return res.status(500).json({
      message: "Internal Server",
      errorMessage: error.message,
      error,
    });
  }
};

exports.deleteSshKey = async (req, res) => {
  try {
    const { id } = req.params;
    const sshKey = await SSHKey.findOneAndUpdate(
      { _id: id },
      { isDeleted: true }
    );
    return res.json({
      message: "SSH Key deleted",
      sshKey,
    });
  } catch (error) {
    return res.status(500).json({
      message: "Internal Server",
      errorMessage: error.message,
      error,
    });
  }
};
