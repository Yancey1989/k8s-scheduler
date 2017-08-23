FROM ubuntu:16.04
ADD ./k8s-scheduler /usr/bin/k8s-scheduler
CMD ["k8s-scheduler"]

