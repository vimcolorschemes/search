# vimcolorschemes/search

This a basic proxy API for elasticsearch used by
[vimcolorschemes](https://github.com/vimcolorschemes/vimcolorschemes).

## Get started

First, fork and clone the project to your machine. Then, choose one of the 2
paths below to run the app.

### Local

  1. [Install elasticsearch locally](https://www.elastic.co/start)
  2. Make sure elasticsearch is running
  3. `cd` into the root of the repository
  2. `npm install`
  3. `npm start`

### Docker

  1. `cd` into the root of the repository
  2. Run `bin/docker-setup`

That's it. 2 docker containers will be built and started, 1 for elasticsearch,
and 1 for the search proxy.

### Communication with the vimcolorschemes app

If the environment variables were not touched, the following happened:

- elasticsearch runs at `http://localhost:9200`
- search proxy runs at `http://localhost:3000`

The elasticsearch instance is called when the vimcolorschemes app builds. All the
data is uploaded to it.

The search proxy is called on the client side, when making a search request
through the website.

## The future of vimcolorschemes/search

The search needs of vimcolorschemes are pretty basic, and they don't require
a search engine as powerful as elasticsearch.

One day, vimcolorschemes/search will become a standard API.

Its job will be the following:

1. Receive data at vimcolorschemes build time
2. Generate a search index (using a library)
3. Store the search index
4. Receive search requests and return the corresponding data
