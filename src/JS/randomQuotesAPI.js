var request = require("request");

var options = {
  method: "GET",
  url: "http://quotes.rest/qod.json",
  qs: { category: "inspire" },
  headers: {
    "Cache-Control": "no-cache"
  }
};

request(options, function(error, response, body) {
  if (error) throw new Error(error);

  console.log(body);
});
