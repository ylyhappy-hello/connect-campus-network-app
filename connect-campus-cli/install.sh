basedir=$(pwd)
# read -r userno
# read -r passwd
# echo "$basedir/main $userno $passwd"
echo "$basedir/main 202114600207 qwefsdd435fgsdH@H@" > run.sh
echo "0 */4 * * * $basedir/run.sh" > cron
crontab cron
chmod +x run.sh
