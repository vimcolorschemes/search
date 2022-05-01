# vimcolorschemes/search

This the search API used by [vimcolorschemes](https://github.com/vimcolorschemes/vimcolorschemes).

It's a AWS Lambda function built with Golang and has 2 functions:

- **Store**: Receive repositories from a daily build of the app, and store it in a
  MongoDB database to be used as a search index.
- **Search**: Receive search parameters and return a list of repositories
  matching the request.

## Get started

Instructions on how to set the search API locally are coming soon...
