import { Config, Frameworks } from 'karma';

export default function(config: Config) {
  config.set({
    frameworks: ['jasmine', '@angular-devkit/build-angular'] as Frameworks[],
    plugins: [
      require('karma-jasmine'),
      require('karma-chrome-launcher'),
      require('@angular-devkit/build-angular/plugins/karma')
    ],
    browsers: ['Chrome'],
    reporters: ['progress'],
    singleRun: false
  });
}
