const express = require("express");
const bcrypt = require("bcrypt");
const { generatePassword } = require("../helper/utils");

const router = express.Router();

const users = [];
router.post("/", (req, res) => {
  const { NIK, Role } = req.body;

  if (!NIK || !Role)
    return res.status(400).json({ error: "Must completed the output " });
  if (!/^\d{16}$/.test(NIK))
    return res.status(400).json({
      error: "The NIK must be exactly 16 digits and contain only numbers",
    });

    
    const checkUser = users.find((x) => x.NIK === NIK);
    // console.log("curreent users", che);
  if (checkUser !== undefined)
    return res.status(401).json({ error: "The NIK is already exist!" });

  const rawPassword = generatePassword();
  console.log("password", rawPassword);

  const password = bcrypt.hashSync(rawPassword, 10);
  const user = { id: users.length + 1, NIK, Role, password: password };
  users.push(user);

  return res.json({ NIK, Role, password });
});

module.exports = { createUserRouter: router, users };
