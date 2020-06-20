
Improve error handling!!

expects an ENV "SECRET"


```
docker build -t jonas27test/jwt-backend .
docker push jonas27test/jwt-backend
docker run -p 8000:8080 -e SECRET='secret' jonas27test/jwt-backend
```