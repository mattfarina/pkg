# As an application, I can find general information about a package.

Applications and packages written in Go have a wide variety of meta-information
surrounding them such as name, an image, a homepage other than the package
location and so forth.

An example of this is Kubernetes. It's name, as displayed is something other than
the packages location. There is a logo, a homepage, and other associated data
that would be useful for an application looking to display information about it.

## Solution
Basic information about an application is captured in a machine parsable manner.
The image is either a relative path within the repository or a full URL to an
outside location.

```json
{
    "name": "Kubernetes",
    "image": "logo.png",
    "description": "Manage a cluster of Linux containers as a single system to accelerate Dev and simplify Ops.",
    "keywords": ["container", "linux", "application"],
    "homepage": "http://kubernetes.io",
    "license": "Apache-2.0"
}
```
