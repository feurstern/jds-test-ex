const express = require("express");
const jwt = require("jsonwebtoken");
const bcrypt = require("bcrypt");
const { users } = require("./createUser");
require("dotenv").config();

const router = express.Router();
const secretKey = process.env.JWT_KEY_SECRET;

router.post("/", (req, res) => {
  const { NIK, password } = req.body;
  const user = users.find((x) => x.NIK === NIK);

  let isExist = false;
  if (user != undefined) isExist = true;
  else return res.status(401).json({ error: "The NIK is not found!" });

//   console.log("current user", user);

//   console.log("password sent:", password);
//   console.log("bycript password ssent", bcrypt.hashSync(password, 10))
//   console.log(
//     "result password validation",
//     !bcrypt.compareSync(password, user.password)
//   );

  if (!bcrypt.compareSync(password, user.password))
    return res.status(401).json({ error: "Invalid credentials" });

  const token = jwt.sign(
    { id: user.id, NIK: user.NIK, Role: user.Role },
    secretKey,
    { expiresIn: "1h" }
  );
  return res.json({ id: user.id, NIK: user.NIK, Role: user.Role, token });
});

module.exports = router;
