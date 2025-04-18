@echo off
echo Setting up Go backend...

echo Copying environment file...
copy .env.example .env

echo Done!
echo.
echo To run the application:
echo 1. Make sure Go is installed
echo 2. Run: go mod download
echo 3. Run: go run main.go
echo.
echo The API will be available at http://localhost:5000
