var ShowRoute = Ember.Route.extend({

  model: function(params) {
    return App.GeoJot.find(params.geojot_id);
  }

});

module.exports = ShowRoute;
