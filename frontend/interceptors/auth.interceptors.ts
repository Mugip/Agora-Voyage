typescript
import { Injectable } from '@angular/core';
import { HttpInterceptor, HttpRequest, HttpHandler, HttpEvent } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from 'path/to/auth.service'; // Replace with the actual path to your auth service

@Injectable()
export class AuthInterceptor implements HttpInterceptor {

  constructor(private authService: AuthService) {}

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    // Clone the original request to add the JWT token to the headers
    const clonedRequest = request.clone({
      setHeaders: {
        Authorization: `Bearer ${this.authService.getToken()}` // Replace with the method to get the JWT token from the auth service
      }
    });

    // Pass the cloned request to the next interceptor or the backend API
    return next.handle(clonedRequest);
  }
}
