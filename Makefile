build:
	cd frontend && npm run build && cd ../
	go build .
	chmod +x fileprotector
	zip -r app.zip fileprotector .env.example assets/ templates/
	rm fileprotector