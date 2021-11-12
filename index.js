require("dotenv").config();

const express = require("express");
const cors = require("cors");
const { createProxyMiddleware } = require("http-proxy-middleware");

const ElasticSearchHelper = require("./helpers/elasticSearch");

const HOST = process.env.ELASTICSEARCH_HOST || "http://localhost:9200";
const USERNAME = process.env.ELASTICSEARCH_USERNAME;
const PASSWORD = process.env.ELASTICSEARCH_PASSWORD;

const WEBSITE_ORIGIN = process.env.WEBSITE_ORIGIN || "http://localhost:8000";

const proxy = createProxyMiddleware({
  target: HOST,
  changeOrigin: true,
  auth: USERNAME && PASSWORD ? `${USERNAME}:${PASSWORD}` : undefined,
  onProxyReq: (proxyReq, req) => {
    const body = ElasticSearchHelper.buildElasticSearchRequestBody(req.body);
    const bodyData = JSON.stringify(body);
    proxyReq.setHeader("Content-Type", "application/json");
    proxyReq.setHeader("Content-Length", Buffer.byteLength(bodyData));
    proxyReq.write(bodyData);
  },
});

const app = express();

app.use(
  cors({
    origin: WEBSITE_ORIGIN,
    optionsSuccessStatus: 200,
  }),
);

app.use(express.json());
app.use("/", proxy);
app.listen(process.env.PORT || 3000);
