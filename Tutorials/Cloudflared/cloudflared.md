# Cloud Flare using tunnel argo

### Installation

[Reference](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/install-and-setup)


###### Download the file first

`wget https://github.com/cloudflare/cloudflared/releases/latest/download/cloudflared-linux-amd64.deb`

###### Extract the file

`sudo dpkg -i cloudflared-linux-amd64.deb`

###### Login in the cloudflared

`cloudflared tunnel login`

> after login it will generate a cert.pem

go here
`cd /home/<yours>/.cloudflared/`

> make sure you have this 3 files inside this folder
> - aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.json
> - cert.pem
> - config.yaml

Inside the config.yaml


```yaml
tunnel: aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee
credentials-file: /home/<yours>/.cloudflared/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee.json

ingress:
        - service: http://35.197.134.241
          originRequest:
          originServerName: mismastermaker.com
```

# RUN

`cloudflare tunnel run`
