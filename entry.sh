#!/bin/ash

# start cron
# /usr/sbin/crond -f -l 8
echo "*/1 * * * * /app/main" >> /etc/crontabs/root
crond -l 2 -f