const jwt = require("jsonwebtoken");

module.exports = () => {
  return (req, res, next) => {
    if (req.headers.authorization) {
      try {
        const token = req.headers.authorization.split(" ")[1];
        const payload = jwt.verify(token, "SeCrEtKeY");
        req.user = payload.id;
        next();
      } catch (error) {
        res.status(401).json({ message: "Unauthorized" });
      }
    } else {
      res.status(401).json({ message: "Unauthorized" });
    }
  };
};
