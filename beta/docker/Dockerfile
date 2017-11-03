FROM google/cloud-sdk:latest

ENV PROJECT_PATH /somwut-j/beta

#ENV HOME /home/default

RUN sed -i -e 's/true/false/' /google-cloud-sdk/lib/googlecloudsdk/core/config.json
RUN gcloud components install app-engine-php
RUN apt-get install php5.5

ADD bashrc /home/default/.bashrc
CMD bash

EXPOSE 8080 8000 39774
WORKDIR $PROJECT_PATH