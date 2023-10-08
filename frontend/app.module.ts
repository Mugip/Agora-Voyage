
import {HTTP_INTERCEPTORS} from @angular/common/http ;
// Add the `AuthInterceptor` to the `providers` array in the `@NgModule` decorator.  
 {
     provide: HTTP_INTERCEPTORS,
     useClass: AuthInterceptor,
     multi: true
 }
