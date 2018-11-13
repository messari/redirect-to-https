
# Redirect to HTTPs

Redirects all traffic to https

### Usage

```bash
$ docker run -e PORT=8080 messari/redirect-to-https
```

### Kubernetes

```yaml
- name: redirect-to-https
  image: messari/redirect-to-https
  env:
    - name: PORT
      value: "8080"
  readinessProbe:
    tcpSocket:
      port: 8080
    initialDelaySeconds: 3
    periodSeconds: 10
    successThreshold: 1
  ports:
    - containerPort: 8080
```
