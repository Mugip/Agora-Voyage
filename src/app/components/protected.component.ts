import { Component, OnInit } from '@angular/core';
import { AuthService } from './authentication.go';

@Component({
  selector: 'app-protected',
  templateUrl: './protected.component.html',
  styleUrls: ['./protected.component.css']
})
export class ProtectedComponent implements OnInit {
  user: any; // Placeholder for user data

  constructor(private authService: AuthService) {}

  ngOnInit(): void {
    this.user = this.authService.getUser(); // Retrieve authenticated user data
    // You can perform additional logic and retrieve data specific to the luxury travel app
  }

  // Add any other necessary methods and functionalities for the protected component
