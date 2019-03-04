# Kubernetes - Part Two

This is the helm section in a multipart series about learning Kubernetes on GCP.
To run this code here are the quick steps.

### To startup and deploy helm:
```bash
git clone https://github.com/tinhtrantrung/gke.git
cd ~/gke/part002/scripts
sh startup.sh
sh add_helm.sh
sh add_redis.sh # add redis
```

### To startup and deploy secure helm:
```bash
git clone https://github.com/tinhtrantrung/gke.git
cd ~/gke/part002/scripts
sh startup.sh
sh add_secure_helm.sh
sh add_secure_redis.sh # add redis
```

### To Teardown:
```bash
cd ~/gke/part002/scripts # if necessary
sh teardown.sh
```