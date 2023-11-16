FROM golang:1.21.4-alpine3.18 
WORKDIR /app

# copying go dependencies into container
COPY ./src/go.mod ./src/go.sum ./src/main.go ./src/.env ./ 
# copying cron config into container
COPY ./entry.sh ./
# COPY ./script.sh ./
# COPY ./crontab.txt ./

# downloading go packages INSIDE container
RUN go mod download

# setting up linux permissions on exec script 

# setting up crontab 

# building go app inside container 
RUN CGO_ENABLED=0 GOOS=linux && go build -o /main 


RUN chmod 755 /app/entry.sh 

RUN echo '*/1  *  *  *  * /main' >> /etc/crontabs/root
# RUN /usr/bin/crontab /app/crontab.txt
# CMD ["tail", "-f", "/dev/null"]
# CMD ["crond","-f", "-l","2" ]
CMD ["/usr/sbin/crond", "-f", "-l", "2"]
# CMD ["/app/entry.sh"]
# CMD ["./main"]