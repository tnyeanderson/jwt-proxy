# jwt-proxy

jwt-proxy is a proxying webserver which validates JWTs along the way.

This project is meant to be an incredibly simple example of how to use the
[NGINX auth request
module](https://nginx.org/en/docs/http/ngx_http_auth_request_module.html) to
create a reverse proxy server which will validate the JWT of any request before
forwarding the request to your actual app. This is usually used to add
authentication to an existing API server without having to modify the API code.

The expected way to use this repo is to fork the project, modify the associated
go program as needed to perform whatever validation you want (right now it only
validates JWTs based on a configured issuer and client ID--this may or may not
be enough for you), build the image, and use it as a sidecar container to your
actual application or API. An example of a kubernetes deployment manifest is
provided.

This project should not be considered "production-ready". It is a barebones
example of using NGINX to add authentication to an existing HTTP backend. You
should definitely do your own research and add any additional security
requirements, validation, and configurations that may be required by your use
case or employer.
