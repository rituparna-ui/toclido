const express = require("express");
const mongoose = require("mongoose");
const cors = require("cors");

const apiRoutes = require("./router");

const app = express();
const port = process.env.PORT || 3000;

app.use(cors());
app.use(express.json());

app.use("/api", apiRoutes);

mongoose
  .connect("mongodb://localhost:27017/toclido")
  .then(() => {
    console.log("MONGODB CONNECTED");
    return app.listen(port);
  })
  .then(() => {
    console.log("EXPRESS STARTED");
  })
  .catch((err) => {
    console.log(err);
    process.exit(1);
  });
