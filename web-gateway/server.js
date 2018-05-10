'use strict';

const express = require('express');
const bodyParser = require('body-parser');
var request = require('request');

const app = express();

app.use(bodyParser.urlencoded({extended: false}));
app.use(bodyParser.json());

const PORT = 8080;
const HOST = '0.0.0.0';

app.get('/status', (req, res) => {
  console.log("/status called.");
  res.status(200).send("OK");
});

app.post('/user', (req, res) => {
  console.log("/user called.");
  console.log("/user req.body : " + req.body);
  
  if(!req || !req.body)
  {
     var errorMsg = "Invalid argument sent";
     console.log(errorMsg);
     return res.status(500).send(errorMsg);
  }

  console.log("calling " + process.env.VCSA_MANAGER);
  const options = {
    url: process.env.VCSA_MANAGER,
    method: 'GET',
    headers: {
      'Accept': 'application/json'
    }
  };

  request(options, function(err, resDoc, body) {
    console.log("callback : " + body);
    if(err)
    {
      console.log("ERROR: " + err);
      return res.send(err);
    }

    console.log("statusCode : " + resDoc.statusCode);
    if(resDoc.statusCode != 200)
    {
      console.log("ERROR code: " + res.statusCode);
      return res.status(500).send(resDoc.statusCode);
    }

    return res.send({"ok" : body});
  });

});


app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
