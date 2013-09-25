var ShowController = Ember.ObjectController.extend({

  message: '',

  send: function(){
    if(!this.conn){
      this.build_connection();
      if(!this.conn){
        this.appendLog($("<div><b>Connection unavailable.</b></div>"));
      }
    }
    if(!this.get('message')){
      return false;
    }
    this.conn.send(this.get('message'));
    this.set('message', '');
  },

  setupController: function(controller, model) {
    this.build_connection();
  },

  build_connection: function(){
    var _self = this;
    var geojot = this.get('model');
    if (window["WebSocket"]) {
      this.conn = new WebSocket("ws://localhost:5000/ws");
      this.conn.onclose = function(evt) {
        _self.conn = null;
        _self.appendLog($("<div><b>Connection closed.</b></div>"));
      };
      this.conn.onmessage = function(evt) {
        _self.appendLog($("<div/>").text(evt.data));
      };
    } else {
      this.appendLog($("<div><b>Your browser does not support WebSockets.</b></div>"));
    }
  },

  appendLog: function(msg) {
    var log = $("#log");
    var d = log[0];
    var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
    msg.appendTo(log);
    if(doScroll){
      d.scrollTop = d.scrollHeight - d.clientHeight;
    }
  }

});

module.exports = ShowController;

