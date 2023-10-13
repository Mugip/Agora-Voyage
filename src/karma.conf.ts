import { ConfigOptions } from 'karma';

module.exports = (config: ConfigOptions) => {
  config.set({
    frameworks: ['jasmine', '@angular-devkit/build-angular'],
    plugins: [
      require('karma-jasmine'),
      require('karma-chrome-launcher'),
      require('@angular-devkit/build-angular/plugins/karma')
    ],
    browsers: ['Chrome'],
    reporters: ['progress', 'kjhtml'],
    singleRun: false
  });
};
