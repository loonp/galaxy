module.exports = function(grunt) {

  // 配置Grunt各种模块的参数
  grunt.initConfig({
    concat: { /* concat的参数 */ },
    uglify: { /* uglify的参数 */ },
    watch:  { /* watch的参数 */ },
    cssmin: {
      minify: {
        expand: true,
        cwd: 'static/css/manage/',
        src: ['*.css', '!*.min.css'],
        dest: 'static/css/manage',
        ext: '.min.css'
      }
    }
  });

grunt.loadNpmTasks('grunt-contrib-cssmin');

grunt.registerTask('default', ['cssmin:minify']);

};