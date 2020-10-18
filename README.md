# vimcolorschemes/search

This a basic proxy API for elasticsearch used by
[vimcolorschemes](https://github.com/reobin/vimcolorschemes).

## Get started

1. Make sure `elasticsearch` is running
2. `npm install`
3. `npm start`

## The future of vimcolorschemes/search

The search needs of vimcolorschemes are pretty basic, and they don't require
a search engine as powerful as elasticsearch.

One day, vimcolorschemes/search will become a standard API.

Its job will be the following:

1. Receive data at vimcolorschemes build time
2. Generate a search index (using a library)
3. Store the search index
4. Receive search requests and return the corresponding data
