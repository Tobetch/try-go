# URL

- [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)

## CURL

- GET(all items)
  - ex: ```curl http://localhost:8080/albums```
- GET(specific item)
  - ex: ```curl http://localhost:8080/albums/2```
- POST
  - ex:
  ```curl http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'```
