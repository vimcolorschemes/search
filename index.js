require("dotenv").config();

const express = require("express");
const cors = require("cors");
const { createProxyMiddleware } = require("http-proxy-middleware");

const proxy = createProxyMiddleware({
  target: process.env.ELASTICSEARCH_HOST || "http://localhost:9200",
  changeOrigin: true,
});

const app = express();
app.use(cors());
app.use("/", proxy);
app.listen(process.env.PORT || 3000);
