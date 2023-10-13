import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  isLoggedIn: boolean;
  username: string;

  constructor() { }

  ngOnInit(): void {
    // Check if the user is logged in
    // and fetch username from local storage or API

    // Example implementation
    this.isLoggedIn = this.checkLoggedInStatus();
    this.username = this.fetchUsername();
  }

  private checkLoggedInStatus(): boolean {
    // Implement your logic here to check if the user is logged in
    // Return true if logged in, else return false
  }

  private fetchUsername(): string {
    // Implement your logic here to fetch the username
    // Return the username as a string
  }

  logout(): void {
    // Implement your logout logic here
    // Clear any stored tokens or user information from local storage
    // Redirect the user to the login page or home page
import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-logout',
  template: `
    <button (click)="logout()">Logout</button>
  `,
})
export class LogoutComponent {
  constructor(private router: Router) {}

  logout(): void {
    // Clear any stored tokens or user information from local storage
    localStorage.removeItem('token');

    // Redirect the user to the login page or home page
    this.router.navigate(['/login']);
  }
  }
  }
} 
