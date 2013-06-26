// require other, dependencies here, ie:
// require('./vendor/moment');

require('../vendor/jquery');
require('../vendor/handlebars');
require('../vendor/ember');
require('../vendor/ember-data'); // delete if you don't want ember-data

var App = Ember.Application.create();
App.Store = require('./store'); // delete if you don't want ember-data

App.Utils = {
  current_location: null,
  CurrentLocation: function(callback){
    if(this.current_location){
      callback(this.current_location);
      return;
    }
    navigator.geolocation.getCurrentPosition(function(location){
      this.current_location = location;
      callback(location);
    }.bind(this));
  }
};

App.deferReadiness();
App.Utils.CurrentLocation(function(location){
  App.set('location', location);
  App.advanceReadiness();
});

module.exports = App;

