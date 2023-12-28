**1 To dump the sample_mflix.movies out of atlas cluster**

`mongodump --uri mongodb+srv://user:password@cluster0.g3aer.mongodb.net/sample_mflix --collection=movies`

**2 To restore the sample_mflix movies locally to atlas instance**

`mongorestore --uri mongodb+srv://localhost:27017/sample_mflix --collection=movies`

**3. Command to build locally**

`podman build -t localhost/moviesapi .`

**4. Command to run locally**

`podman run -d -p 8080:8080 --name moviesapi localhost/moviesapi`

**5. Command to build to deploy on GCP cloud run**

`podman buildx build --platform linux/amd64 -t moviesapi .`

**6. Command to tag instance**

`podman tag localhost/moviesapi:latest docker.io/eugenetan0/moviesapi:latest`

**7. Command to run on push to Dockerhub**

`podman push docker.io/eugenetan0/moviesapi:latest`

**OR**

Skip steps 5 - 7 and do:

`terraform init`

and

`terraform apply`

to push build, push and deploy directly to GCP cloud run
