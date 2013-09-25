module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    sass: {
      dist: {
        files: [{
          expand: true,
          cwd: 'css',
          src: ['*.scss'],
          dest: './public/css',
          ext: '.css'
        }]
      }
    },

    shell: {
      ember: {
        options: {
          stdout: true
        },
        command: 'ember build'
      }
    },

    watch: {
      files: ['public/js/*.js', 'css/**/*.scss', '!public/js/application.js', '!public/js/index.js', '!public/js/templates.js'],
      tasks: ['default']
    }

  });

  // Load libs
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-shell');

  // Register the default tasks
  grunt.registerTask('default', ['sass', 'shell']);

};