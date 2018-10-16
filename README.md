Michael Lloyd Lee CV Redirector
===============================

Wat?
---
This provides static links to my copies of mv CV at various points in time. My CV is hosted on 
Google Docs or Word Online.  This simply minifies and nicifies the URL.

How?
---
This redirection service makes use of the following technologies:

* A tiny GoLang service to handle redirects.
* Docker to build and execute the GoLang service.
* [zeit.co NOW](https://zeit.co/) to host this as a serverless component.
* GitHub/NOW integration to perform continuous deployments. 

To build the project locally via Docker:

```bash
$ docker build . -t mycv
$ docker run mycv
```

To force a deployment of the project:

```bash
$ now --public
$ now alias cv.michael-lloyd-lee.me.uk
```

Continuous
----------

CI/CD is handled via now [now](https://zeit.co/). The `now.json` ensures that every push to master is deployed to [https://cv.michael-lloyd-lee.me.uk](https://cv.michael-lloyd-lee.me.uk)
