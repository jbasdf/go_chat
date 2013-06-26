(function() {

  var App = window.App = Ember.Application.create();
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

  // Order matters. You have to register types before creating the store
  DS.RESTAdapter.registerTransform('coordinatePoint', {
    serialize: function(serialized) {
      return Ember.isNone(serialized) ? {} : serialized;
    },
    deserialize: function(deserialized) {
      return Ember.isNone(deserialized) ? {} : deserialized;
    }
  });

  DS.Store.create({
    revision: 12,
    adapter: DS.RESTAdapter.create({
      namespace: 'api'
    })
  });

  App.Router.map(function(){
    this.route('create');
    this.route('edit', {path: '/edit/:geojot_id'});
  });

  App.GeoJot = DS.Model.extend({
    name: DS.attr('string'),
    location: DS.attr('coordinatePoint')
  });

  App.IndexRoute = Ember.Route.extend({
    model: function(){
      var jots = Ember.ArrayController.create();
      App.Utils.CurrentLocation(function(location){
        jots.set('content', App.GeoJot.find({lat: location.coords.latitude, lon: location.coords.longitude}));
      });
      return jots;
    },

    events: {
      selectGeoJot: function(geojot){

      },
      deleteGeoJot: function(geojot) {
        geojot.deleteRecord();
        geojot.save();
      }
    }
  });

  App.CreateController = Ember.Controller.extend({
    name: null,
    save: function(){
      App.Utils.CurrentLocation(function(location){
        var geojot = App.GeoJot.createRecord({
          name: this.get('name'),
          location: {
            lat: location.coords.latitude,
            lon: location.coords.longitude
          }
        });
        geojot.save().then(function() {
          this.transitionToRoute('index');
          this.set('name', '');
        }.bind(this));
      }.bind(this));
    }
  });

  App.EditController = Ember.ObjectController.extend({
    save: function() {
      var geojot = this.get('model');
      geojot.save().then(function() {
        this.transitionToRoute('index');
      }.bind(this));
    }
  });

})();