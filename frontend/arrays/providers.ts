typescript
import { NgModule } from '@angular/core';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { AuthInterceptor } from 'Agora-voyage/frontend/interceptors/auth.interceptor.ts'; // Replace with the actual path to your auth interceptor

@NgModule({
  imports: [
    HttpClientModule
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    }
  ]
})
export class AppModule { }
