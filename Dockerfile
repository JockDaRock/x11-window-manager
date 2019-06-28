# FROM balenalib/intel-nuc-debian:stretch-run
FROM ubuntu

ARG DEBIAN_FRONTEND=noninteractive

# Install XORG
#RUN install_packages xserver-xorg=1:7.7+19 \
RUN apt-get update && apt-get -y install xserver-xorg \
  xserver-xorg-input-evdev \
  xinit \
  xfce4 \
  xfce4-terminal \
  x11-xserver-utils \
  dbus-x11 \
  matchbox-keyboard \
  xterm \
  chromium-bsu \
  apt-transport-https \
  curl

RUN curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg \
    && install -o root -g root -m 644 microsoft.gpg /etc/apt/trusted.gpg.d/ \
    && sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list' \
    #&& apt-get install apt-transport-https \
    && apt-get update \
    && apt-get install code

# Disable screen from turning it off
RUN echo "#!/bin/bash" > /etc/X11/xinit/xserverrc \
  && echo "" >> /etc/X11/xinit/xserverrc \
  && echo 'exec /usr/bin/X -s 0 dpms' >> /etc/X11/xinit/xserverrc 

# Setting working directory
WORKDIR /usr/src/app

COPY . ./

ENV UDEV=1

# Avoid requesting XFCE4 question on X start
ENV XFCE_PANEL_MIGRATE_DEFAULT=1

CMD ["bash", "start_x86.sh"]