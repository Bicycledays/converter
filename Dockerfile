FROM debian:10

WORKDIR /root

RUN apt-get update
RUN apt-get install -y --no-install-recommends wget
RUN libxinerama1 libdbus-1-3 libglib2.0-0
RUN libcups2 libcairo2 libsm6 default-jre libreoffice-java-common mc
RUN rm -rf /var/lib/apt/lists/*

RUN wget --no-check-certificate https://download.documentfoundation.org/libreoffice/stable/7.3.1/deb/x86_64/LibreOffice_7.3.1_Linux_x86-64_deb.tar.gz
RUN tar -zxvf LibreOffice_7.3.1_Linux_x86-64_deb.tar.gz
RUN dpkg -i LibreOffice_7.3.1.3_Linux_x86-64_deb/DEBS/*.deb

COPY . /app
EXPOSE 8000

CMD ["/app/main"]