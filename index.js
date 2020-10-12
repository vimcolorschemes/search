require("dotenv").config();

const express = require("express");
const cors = require("cors");
const { createProxyMiddleware } = require("http-proxy-middleware");

const HOST = process.env.ELASTICSEARCH_HOST || "http://localhost:9200";
const USERNAME = process.env.ELASTICSEARCH_USERNAME;
const PASSWORD = process.env.ELASTICSEARCH_PASSWORD;

const proxy = createProxyMiddleware({
  target: HOST,
  changeOrigin: true,
  auth: USERNAME && PASSWORD ? `${USERNAME}:${PASSWORD}` : undefined,
});

const app = express();
app.use(cors());
app.use("/", proxy);
app.listen(process.env.PORT || 3000);
