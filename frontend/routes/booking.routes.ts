typescript
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BookingService {
  baseUrl = 'http://localhost:8000'; // Base URL of the backend server

  constructor(private http: HttpClient) {}

  createBooking(booking: any): Observable<any> {
    const url = `${this.baseUrl}/bookings`;
    return this.http.post<any>(url, booking);
  }

  getBookings(userId: string): Observable<any[]> {
    const url = `${this.baseUrl}/users/${userId}/bookings`;
    return this.http.get<any[]>(url);
  }

  getBooking(bookingId: string): Observable<any> {
    const url = `${this.baseUrl}/bookings/${bookingId}`;
    return this.http.get<any>(url);
  }
}
