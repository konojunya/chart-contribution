deploy:
	npm install
	dep ensure
	npm run build
	git push heroku master
