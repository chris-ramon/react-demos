require('node-jsx').install();
var system = require('system');
var args = system.args;
var name = args[2];
var React = require('react');
var HelloMessage = require('./hellomessage.js');
var html = React.renderToString(React.createElement(HelloMessage, {name: name}));
console.log(html);
