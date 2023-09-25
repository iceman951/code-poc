const express = require('express');
const router = express.Router();
const userController = require("../controllers/users");

/* GET users listing. */
router.get("/", userController.show);

module.exports = router;
