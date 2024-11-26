const jwt = require("jsonwebtoken");
const CliToken = require("./../models/cli-token");

module.exports = () => {
  return async (req, res, next) => {
    if (req.headers.authorization) {
      const [type, token] = req.headers.authorization.split(" ");
      if (type === "Bearer") {
        try {
          const payload = jwt.verify(token, "SeCrEtKeY");
          req.user = payload.id;
          return next();
        } catch (error) {
          return res.status(401).json({ message: "Unauthorized" });
        }
      }
      if (type === "Cli") {
        const cliToken = await CliToken.findOne({ token: token }).populate(
          "user",
          "_id"
        );
        if (cliToken) {
          req.user = cliToken.user._id;
          return next();
        } else {
          return res.status(401).json({ message: "Unauthorized" });
        }
      }
      return res.status(401).json({ message: "Unauthorized" });
    } else {
      return res.status(401).json({ message: "Unauthorized" });
    }
  };
};
