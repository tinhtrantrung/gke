# Kubernetes - Part One

This is the first section in a multipart series about learning Kubernetes on GCP.
To run this code here are the quick steps.

### To startup and deploy:
```bash
git clone https://github.com/tinhtrantrung/gke.git
cd gke/part001/scripts
sh startup.sh
sh deploy.sh
sh check-endpoint.sh
```

### To Teardown:
```bash
cd gke/part001/scripts # if necessary
sh teardown.sh
```