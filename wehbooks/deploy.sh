
# Logging
#di lihat terlebih dahulu untuk lokasi foldernya
LOG_FILE="/home/kelas-santai/webhooks/webhook-automation/wehbooksdeploy.log"
echo "$(date) - Starting deployment" >> $LOG_FILE

#masuk ke projeck yang akan di tarik
cd /home/kelas-santai/pertemuan-1/bagja/myapp/backend-golang

# 1. HAPUS SUDO DARI GIT PULL
git pull origin main >> $LOG_FILE 2>&1

# 3. JALANKAN DOCKER TANPA SUDO
docker compose down >> $LOG_FILE 2>&1
docker compose up -d --build >> $LOG_FILE 2>&1

echo "$(date) - Deployment completed" >> $LOG_FILE

# Nonaktifkan venv
