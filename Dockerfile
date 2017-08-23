FROM ubuntu:16.04
ADD ./k8s-scheduler /usr/bin/k8s-scheduler
#ADD ./run.sh /usr/bin/run.sh
#ADD ./kubectl /usr/bin/kubectl
#RUN apt-get update -y && apt-get install -y jq curl
CMD ["k8s-scheduler"]

