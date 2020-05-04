.PHONY: go-gin nodejs-express python-flask

go-gin:
	cd go-gin && go run server.go && cd .. 

nodejs-express:
	cd nodejs-express && npm start && cd ..

python-flask:
	cd python-flask && python3 server.py && cd ..