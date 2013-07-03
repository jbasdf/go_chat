var App = require('./app');

App.Router.map(function() {
  this.route('create');
  this.route('edit', {path: '/edit/:geojot_id'});
  this.route('show', {path: '/:geojot_id'});
});

