# hibergate

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

Now that you have been warned...

The main use case for this tool is to allow TCP services to scale to zero when idling. Just as an example, imagine the
following scenario: You have a kubernetes cluster running in one of the major clouds. You have cluster auto-scaling all
set up so that kubernetes nodes are automatically added to or removed from your cluster based on the running pods'
resource requirements.

Now you want to run applications in this cluster that are idle most of the time. When there are no active TCP connections,
the applications' pods don't need to be running and consuming (costly) cloud resources. That's where socklaunchd comes in.

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
