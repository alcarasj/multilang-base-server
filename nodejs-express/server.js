const express = require('express');
const PORT = 8080;

var server = express();

server.get("/", (req, res) => {
	res.send("Welcome! This is a base server written in NodeJS using the Express framework.");	
});

server.listen(PORT, (error) => {
	if (error) {
		return console.log("Server failed to start.", err);
	}
	console.log("Server listening on port " + PORT + ".");
});
