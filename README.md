# hibergate

The main use case for this tool is to allow TCP services to scale to zero when idling. Just as an example, imagine the
following scenario: You have a kubernetes cluster running in one of the major clouds. You have cluster auto-scaling all
set up so that kubernetes nodes are automatically added to or removed from your cluster based on the running pods'
resource requirements.

Now you want to run applications in this cluster that are idle most of the time. When there are no active TCP connections,
the applications' pods don't need to be running and consuming (costly) cloud resources. That's where hibergate comes in.

First of all, there are similar projects that offer some kind of scale-to-zero support. These include, but are not limited
to, the following:
* Osiris scale-to-zero component for Kubernetes - https://github.com/deislabs/osiris
* Components of Knative - https://knative.dev/
* Nginx with the lua scripting extensions and some custom scripting - https://www.nginx.com/resources/wiki/modules/lua/
and specifically the OpenResty project https://openresty.org/en/
* systemd socket-based activation - https://www.freedesktop.org/software/systemd/man/systemd.socket.html
* the diverse inetd flavours

The reason why I still wrote hibergate is that all of the existing solutions I had tried were quite complex and/or
limited to certain use cases. This bugged me especially because it seemed so simple to just determine when some TCP
service is idle, and then running some command(s) to stop and launch the service accordingly.

The actual documentation for this project will follow soon. Stay tuned. :)

Copyright 2020 Michael Werner

Licensed under the Apache License, Version 2.0 (the "License");
you may not use these files except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.