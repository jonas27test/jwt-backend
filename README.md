
Improve error handling!!

expects an ENV "SECRET"

## krakend

docker run -p 8080:8080 -v "${PWD}/krakend.json:/etc/krakend/krakend.json" devopsfaith/krakend run -d -c /etc/krakend/krakend.json