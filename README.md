# k8s-scheduler
A implement of Kubernetes Scheduler written by Go,just assign pod to node randomly.

## How to run 
- Append `kubernetes <k8s master host>` to your host file `/etc/hosts`
- Add user key files under `./test` folder, like 
    ```bash
        ./test/
        ./test//admin-key.pem
        ./test//admin.pem
        ./test//ca.pem
    ```
- Execute `go run main.go`