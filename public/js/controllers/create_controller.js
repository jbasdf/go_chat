var CreateController = Ember.Controller.extend({
  name: null,
  save: function(){

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

  }
});

module.exports = CreateController;

