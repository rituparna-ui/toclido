const sshpk = require("sshpk");

const SSHKey = require("./../models/ssh-key");

exports.createSshKey = async (req, res) => {
  try {
    const { name, publicKey } = req.body;
    if (!name || !publicKey) {
      return res.status(400).json({ message: "Please provide all the fields" });
    }

    let key;
    try {
      key = sshpk.parseKey(publicKey, "ssh");
    } catch (error) {
      return res.status(400).json({ message: "Invalid public key" });
    }

    const existingKey = await SSHKey.findOne({ publicKey });
    if (existingKey) {
      return res.status(400).json({ message: "Key already exists" });
    }

    await SSHKey.create({
      name,
      publicKey,
      user: req.user,
    });

    return res.json({
      message: "SSH Key Added",
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
        fingerprint: parsedKey.fingerprint("sha256").toString("hex"),
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
