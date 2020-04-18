FROM ubuntu:18.04
RUN apt upgrade -y && \
	apt update -y && \
	apt install golang socat git -y

RUN go get github.com/GeistInDerSH/Covid19-Watch/covid_data

COPY . .

CMD ["/bin/bash", "run.sh"]
