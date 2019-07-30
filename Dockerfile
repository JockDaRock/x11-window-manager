FROM balenalib/intel-nuc-ubuntu:bionic
#FROM balenalib/intel-nuc-ubuntu

ARG DEBIAN_FRONTEND=noninteractive

# Install XORG
#RUN install_packages xserver-xorg=1:7.7+19 \
RUN apt-get update && apt-get -y install \
  xserver-xorg \
  xserver-xorg-input-all \
  xinit \
  #xfce4 \
  #xfce4-terminal \
  x11-xserver-utils \
  #ubuntu-desktop \
  dbus-x11 \
  matchbox-keyboard \
  xterm \
  apt-transport-https \
  curl
  #tasksel

#RUN tasksel install ubuntu-desktop

RUN apt install -y gnome-shell ubuntu-gnome-desktop

RUN apt-get -y install xfce4 xfce4-terminal

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


COPY .dmrc /data/.dmrc
RUN chmod 644 /data/.dmrc

# Setting working directory
WORKDIR /usr/src/app

COPY . ./

ENV UDEV=1

# Avoid requesting XFCE4 question on X start
ENV XFCE_PANEL_MIGRATE_DEFAULT=1

CMD ["bash", "start_x86.sh"]