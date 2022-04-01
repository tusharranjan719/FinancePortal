## Run

Starts `billsplit` in the background.

```sh
docker-compose up --build -d
```

## Destroy

Removes all docker-related `billsplit` resources, including the database's volume.

```sh
docker-compose down -v
```
