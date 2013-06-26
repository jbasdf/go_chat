// Order matters. You have to register types before creating the store
DS.RESTAdapter.registerTransform('coordinatePoint', {
  serialize: function(serialized) {
    return Ember.isNone(serialized) ? {} : serialized;
  },
  deserialize: function(deserialized) {
    return Ember.isNone(deserialized) ? {} : deserialized;
  }
});

module.exports = DS.Store.extend({
  revision: 12,
  adapter: DS.RESTAdapter.create({
    namespace: 'api'
  })
});

