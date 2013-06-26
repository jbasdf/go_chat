var GeoJot = DS.Model.extend({

  name: DS.attr('string'),

  location: DS.attr('coordinatePoint')

});

module.exports = GeoJot;

