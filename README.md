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


# ==============================================================================
### API Document

http://localhost:8088/api/v1/swagger/index.html#/


# ==============================================================================
# Docker support

Monitor live logs with docker compose

```bash
docker compose logs -f app
```