import { Config } from 'karma';
module.exports = (config: Config) => {
  config.set({
    basePath: '',
    frameworks: ['jasmine', '@angular-devkit/build-angular'],
    plugins: [
      require('karma-jasmine'),
      require('karma-chrome-launcher'),
      require('karma-coverage'),
      require('@angular-devkit/build-angular/plugins/karma')
    ],
    client: {
      clearContext: false // leave Jasmine Spec Runner output visible in browser
    },
    coverageReporter: {
      dir: require('path').join(__dirname, 'coverage'),
      reporters: [
        { type: 'html', subdir: 'report-html' },
        { type: 'lcovonly', subdir: 'report-lcov' },
        { type: 'text-summary' }
      ]
    },
    reporters: config.angularCli && config.angularCli.codeCoverage
      ? ['progress', 'coverage']
      : ['progress', 'kjhtml'],
    port: 9876,
    colors: true,
    logLevel: config.LOG_INFO,
    autoWatch: true,
    browsers: ['Chrome'],
    singleRun: false,
    restartOnFileChange: true
  });
}
