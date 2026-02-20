const express = require("express");
const app = express();

app.get("/hello", (req, res) => {
  const channel = req.header("X-CHANNEL") || "NOT RECEIVED";
  const prefix = req.header("X-Forwarded-Prefix") || "NONE";

  res.json({
    message: "Dummy backend reached!",
    channelReceived: channel,
    prefixReceived: prefix
  });
});

const PORT = process.env.PORT || 3001;
app.listen(PORT, () => console.log("Dummy backend running on port " + PORT));