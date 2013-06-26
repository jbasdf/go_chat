var EditController = Ember.ObjectController.extend({
  save: function() {
    var geojot = this.get('model');
    geojot.save().then(function() {
      this.transitionToRoute('index');
    }.bind(this));
  }
});

module.exports = EditController;

