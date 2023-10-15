declare function describe(description: string, specDefinitions: () => void): void;
declare function it(expectation: string, assertion: () => void): void;
declare function expect(actual: any): jasmine.Matchers;
declare function beforeEach(action: () => void): void;
declare function afterEach(action: () => void): void;
declare function spyOn(object: any, method: string): jasmine.Spy;
declare function runs(asyncMethod: Function): void;
declare function waitsFor(latchFunction: () => boolean, failureMessage?: string, timeout?: number): void
