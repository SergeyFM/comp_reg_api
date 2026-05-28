@echo off

echo ==========================================
echo Building Docker image...
echo ==========================================

docker build -f ../docker_compose/Dockerfile -t comp-reg-api:0.1.0 ..

IF %ERRORLEVEL% NEQ 0 (
    echo.
    echo Docker build failed.
    pause
    exit /b 1
)

echo.
echo ==========================================
echo Saving Docker image to TAR...
echo ==========================================

docker save comp-reg-api:0.1.0 -o comp-reg-api_0.1.0.tar

IF %ERRORLEVEL% NEQ 0 (
    echo.
    echo Docker save failed.
    pause
    exit /b 1
)

echo.
echo ==========================================
echo SUCCESS
echo ==========================================
echo TAR file created:
echo.
echo   comp-reg-api_0.1.0.tar
echo.

pause