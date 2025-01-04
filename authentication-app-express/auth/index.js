const express = require("express");
const {createUserRouter} = require("./createUser");
const loginRouter = require("./login");
const claimsRouter = require("./claim");

const router = express.Router();

router.use("/create", createUserRouter);
router.use("/login", loginRouter);
router.use("/private", claimsRouter);

module.exports = router;
