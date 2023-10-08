typescript
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private apiUrl = 'http://localhost:8080'; // Replace with your actual API endpoint

  constructor(private http: HttpClient) { }

  login(username: string, password: string): Observable<string> {
    const authUrl = `${this.apiUrl}/login`;

    return this.http.post<string>(authUrl, { username, password }, { responseType: 'text' as 'json' });
  }
}
import { Component } from '@angular/core';
import { AuthService } from './authservices';

@Component({
  selector: 'app-login',
  template: `
    <form (ngSubmit)="onLogin()">
      <input type="text" [(ngModel)]="username" placeholder="Username">
      <input type="password" [(ngModel)]="password" placeholder="Password">
      <button type="submit">Login</button>
    </form>
  `
})
export class LoginComponent {
  username: string = '';
  password: string = '';

  constructor(private authService: AuthService) { }

  onLogin(): void {
    this.authService.login(this.username, this.password).subscribe(
      token => {
        // Store the token or perform any necessary actions
        console.log('Token:', token);
      },
      error => {
        // Handle the login error
        console.error('Login error:', error);
      }
    );
  }
      }
