const express = require("express");
const request = require("request");
const cors = require("cors");

require("dotenv").config();

const elasticsearchHost =
  process.env.ELASTICSEARCH_URL || "http://localhost:9200";

const port = Number(process.env.ELASTICSEARCH_API_PORT || 3000);

const app = express();

app.use(cors());

app.use("/", (req, res) => {
  if (req.url !== "/favicon.ico") {
    if (!(req.method == "GET" || req.method == "POST")) {
      errMethod = {
        error:
          req.method + " request method is not supported. Use GET or POST.",
      };
      res.write(JSON.stringify(errMethod));
      res.end();
      return;
    }

    const url = `${elasticsearchHost}${req.url}`;
    req
      .pipe(
        request({
          uri: url,
          auth: {
            user: "username",
            pass: "password",
          },
          headers: {
            "accept-encoding": "none",
          },
          rejectUnauthorized: false,
        }),
      )
      .pipe(res);
  }
});

app.listen(port, () => {
  console.log(`elasticsearch is running on ${elasticsearchHost}`);
  console.log(
    `vimcolorschemes-search API is running on http://localhost:${port}`,
  );
});
