@echo off

echo ==========================================
echo Stopping and removing old container...
echo ==========================================

docker stop comp-reg-api >nul 2>&1
docker rm comp-reg-api >nul 2>&1

echo.
echo ==========================================
echo Removing old image...
echo ==========================================

docker rmi comp-reg-api:0.1.0 >nul 2>&1

echo.
echo ==========================================
echo Loading image from TAR...
echo ==========================================

docker load -i comp-reg-api_0.1.0.tar

IF %ERRORLEVEL% NEQ 0 (
    echo.
    echo Failed to load image.
    pause
    exit /b 1
)

echo.
echo ==========================================
echo Starting container...
echo ==========================================

start "" docker run ^
  --name comp-reg-api ^
  -p 8080:8080 ^
  comp-reg-api:0.1.0

timeout /t 3 >nul

echo.
echo ==========================================
echo Test URLs:
echo ==========================================
echo.
echo http://localhost:8080/health
echo http://localhost:8080/status
echo http://localhost:8080/api/v1/companies
echo.
echo Press any key to stop and remove container...
pause >nul

echo.
echo ==========================================
echo Stopping container...
echo ==========================================

docker stop comp-reg-api
docker rm comp-reg-api

echo.
echo ==========================================
echo Done.
echo ==========================================

pause