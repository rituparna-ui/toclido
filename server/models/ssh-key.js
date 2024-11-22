const mongoose = require("mongoose");

const sshKeySchema = new mongoose.Schema(
  {
    name: {
      type: String,
      required: true,
    },
    publicKey: {
      type: String,
      required: true,
    },
    fingerprint: {
      type: String,
      required: true,
    },
    user: {
      type: mongoose.Schema.Types.ObjectId,
      ref: "User",
      required: true,
    },
    isDeleted: {
      type: Boolean,
      default: false,
    },
    lastUsedAt: {
      type: Date,
      default: Date.now,
    },
  },
  {
    timestamps: true,
  }
);

module.exports = mongoose.model("SshKey", sshKeySchema);
