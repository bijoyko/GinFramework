FROM golang:latest

RUN mkdir -p G/BuildsForGithub/GinFramework/GinGorm

WORKDIR -p G/BuildsForGithub/GinFramework/GinGorm

COPY . G/BuildsForGithub/GinFramework/GinGorm

RUN go install GinGorm

CMD /go/bin/GinGorm

EXPOSE 8080