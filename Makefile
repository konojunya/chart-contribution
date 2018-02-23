install:
	npm install

build:
	npm run build

serve:
	env PORT=8000 go run main.go

deploy:
	npm install
	dep ensure
	npm run build
	git push heroku master