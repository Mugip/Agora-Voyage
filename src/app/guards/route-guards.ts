import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';
import { AuthService } from './auth.service'; 

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private router: Router) {}

  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
    // Check for authentication logic here
    // For example, you can use a service to check if the user is authenticated
@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private authService: AuthService, private router: Router) {}

  canActivate(): boolean {
    if (this.authService.isAuthenticated()) {
      return true;
    } else {
      this.router.navigate(['/login']);
      return false;
    }
  }
}

    // Replace this with your actual authentication check
    const isAuthenticated = true;

    if (isAuthenticated) {
      return true; // Allow navigation to the requested route
    } else {
      this.router.navigate(['/login']); // Redirect to the login page if not authenticated
      return false; // Block navigation to the requested route
    }
  }
}

export class RoleGuard implements CanActivate {
  constructor(private router: Router) {}

  canActivate(next: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
    // Check for role-based authorization logic here
    // For example, you can use a service to check if the user has the required role
    @Injectable()
export class RoleGuard implements CanActivate {
  constructor(private authService: AuthService, private router: Router) {}

  canActivate(): boolean {
    if (this.authService.hasRequiredRole('admin')) {
      return true;
    } else {
      this.router.navigate(['/unauthorized']);
      return false;
    }
  }
}

    // Replace this with your actual role-based authorization check
    const hasRequiredRole = true;

    if (hasRequiredRole) {
      return true; // Allow navigation to the requested route
    } else {
      this.router.navigate(['/unauthorized']); // Redirect to unauthorized page if role is not allowed
      return false; // Block navigation to the requested route
    }
  }
              }
