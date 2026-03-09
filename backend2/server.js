const express = require("express");
const app = express();

app.get("/hello", (req, res) => {
  const requestId = req.header("X-Request-ID") || "NOT RECEIVED";
  const service   = req.header("X-Service") || "NONE";

  res.json({
    message:           "Dummy backend 2 reached!",
    requestIdReceived: requestId,
    serviceReceived:   service
  });
});

const PORT = process.env.PORT || 3002;
app.listen(PORT, () => console.log("Dummy backend 2 running on port " + PORT));
