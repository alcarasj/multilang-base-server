import express from "express";
const PORT = 8080;

const server: express.Application = express();

server.get("/", (req, res) => {
	res.send("Welcome! This is a base server written in NodeJS using the Express framework.");
});

server.listen(PORT, (error) => {
	if (error) {
		return console.error("Server failed to start.", error);
	}
	console.log("Server listening on port " + PORT + ".");
});