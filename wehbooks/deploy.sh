
# Logging
LOG_FILE="/home/kelas-santai/webhooks/webhooks/deploy.log"
echo "$(date) - Starting deployment" >> $LOG_FILE

# Update code & rebuild
cd /home/kelas-santai/myapps/myapp

# 1. HAPUS SUDO DARI GIT PULL
git pull origin main >> $LOG_FILE 2>&1

# 3. JALANKAN DOCKER TANPA SUDO
docker compose down >> $LOG_FILE 2>&1
docker compose up -d --build >> $LOG_FILE 2>&1

echo "$(date) - Deployment completed" >> $LOG_FILE

# Nonaktifkan venv
