var IndexRoute = Ember.Route.extend({
  events: {
    deleteGeoJot: function(geojot) {
      geojot.deleteRecord();
      geojot.save();
    }
  }
});

module.exports = IndexRoute;

