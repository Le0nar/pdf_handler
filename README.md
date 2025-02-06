# Run MinIO

docker run -p 9000:9000 -p 9001:9001 \
  -e "MINIO_ACCESS_KEY=youraccesskey" \
  -e "MINIO_SECRET_KEY=yoursecretkey" \
  --name minio \
  minio/minio server /data

