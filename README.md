# k8s-scheduler
A implement of Kubernetes Scheduler written by Go,just assign pod to node randomly.

## How to run 
- Append `kubernetes <k8s master host>` to your host file `/etc/hosts`
- Add user key files under `./test` folder, like 
```bash
    ./test/
    ./test/admin-key.pem
    ./test/admin.pem
    ./test/ca.pem
```
- Execute `go run main.go`
- Create a pod with `schedulerName: my-scheduler`, execute `kubectl create -f nginx.yaml`, and you will see the logs like:
```bash
node0
node1
node2
...
There ara 6 nodes in the cluster
There are 76 pods in the cluster
Start assign annotation-second-scheduler
Assign annotation-second-scheduler to node0

```