const express = require("express");
const jwt = require("jsonwebtoken");
require("dotenv").config();
const router = express.Router();

const secretKey = process.env.JWT_KEY_SECRET;

router.get("/", (req, res) => {
  const authHeader = req.headers.authorization;
  console.log("Authorization Header:", authHeader);

  if (!authHeader)
    return res.status(401).json({ error: "Authorization header is required!" });

  const token = authHeader.split(" ")[1];

  if (!token) return res.status(401).json({ error: "Token is required!" });

  try {
    const decode = jwt.verify(token, secretKey);
    return res.json({ data: decode });
  } catch (err) {
    return res.status(401).json({ error: `Invalid token: ${err.message}` });
  }
});

module.exports = router;
