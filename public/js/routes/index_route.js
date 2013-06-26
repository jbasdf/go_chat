var IndexRoute = Ember.Route.extend({
  
  events: {
    selectGeoJot: function(geojot){

    },
    deleteGeoJot: function(geojot) {
      geojot.deleteRecord();
      geojot.save();
    }
  }
});

module.exports = IndexRoute;

