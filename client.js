'use strict';

var net = require('net');

var client = net.connect({port: 11414}, function() {
  client.write(JSON.stringify({
    method: 'Llama.Latency',
    params: [{N: 500}],
    id: 0
  }));
});

client.on('data', function(data) {
  console.log(data.toString());

  var resp = JSON.parse(data);
  if (resp.result === 'ok') {
    client.end();
  }
});

client.on('error', function(err) {
  console.log(err);
});

client.on('end', function() {
  console.log('disconnected from the agent');
});
