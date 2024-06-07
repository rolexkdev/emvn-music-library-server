# emvn-music-library-server
#### Prerequisites 
- Docker and Docker compose

#### How to run project
1. Clone the repository to locally
2. Move to project folder
```shell 
cd ./emvn-music-library-server
```
3. Setup the sample environment variables:

```bash
cp .env.example .env
```
4. Docker compose run without app:
```bash
docker compose up -d
```
5. Finally, run the app on port 8088

http://localhost:8088

6. Can restore database with file `backup.tar.gz`


# API Document
-  Swagger document for api endpoint
http://localhost:8088/api/v1/swagger/index.html#/

## API Endpoint

1. `/uploads`
- API upload file mp3 for tracks or image of album_cover for playlist. Return file_url. You can get file_url to create tracks and playlists.
- FileURL format: http://localhost:8088/api/v1/uploads/{filename}
- Example: http://localhost:8088/api/v1/uploads/NangTho.mp3

2. `/tracks`
API CRUD for tracks

- Example Create a Track
```shell
curl --location 'http://localhost:8088/api/v1/tracks' \
--header 'Content-Type: application/json' \
--data '{
  "album": "ablum1",
  "artist_id": "id123456",
  "duration": 300000,
  "file_url": "http://localhost:8088/api/v1/uploads/NangTho.mp3",
  "genre": "pop",
  "name": "nang tho",
  "release_date": 1717786298000,
  "title": "Nang Tho - Hoang Dung"
}'
```

3. `/playlists`
API CRUD for playlists

4. `/search`
API Search tracks and playlists


# Docker support

Monitor live logs with docker compose

```bash
docker compose logs -f app
```

//TODO
Develop a crawler script and crawl at least 100 playlists from Spotify or YouTube Music. The crawled data needs to be included in the database dump.
Important: Do not use any external library or tool, all data should be crawled from your code.
