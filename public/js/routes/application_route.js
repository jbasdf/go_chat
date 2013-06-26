var ApplicationRoute = Ember.Route.extend({

  jots: Ember.ArrayProxy.create({content: []}),

  model: function() {
    var location = App.get('location');
    var jots = App.GeoJot.find({lat: location.coords.latitude, lon: location.coords.longitude});
    jots.on('didLoad', function() {
      jots.forEach(function(jot) {
        this.jots.pushObject(jot);
      }.bind(this));
    }.bind(this));
    return this.jots;
  }

});

module.exports = ApplicationRoute;

