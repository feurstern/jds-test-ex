const express = require("express");
const bodyParser = require("body-parser");
const authRoutes = require("./auth");
require("dotenv").config();

const app = express();
app.use(bodyParser.json());

app.use("/auth", authRoutes);

module.exports = app;

if (require.main === module) {
  const PORT = process.env.PORT || 3000;
  app.listen(PORT, () => {
    console.log(`Server running on http://localhost:${PORT}`);
  });
}
