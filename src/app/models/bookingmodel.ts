export interface Booking {
  id: number;
  user_id: number;
  start_time: Date;
  end_time: Date;
}

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BookingService {
  private apiUrl = 'http://your-api-url/bookings';

  constructor(private http: HttpClient) { }

  createBooking(userID: number, startTime: Date, endTime: Date): Observable<Booking> {
    const booking: Booking = {
      user_id: userID,
      start_time: startTime,
      end_time: endTime
    };
    return this.http.post<Booking>(this.apiUrl, booking);
  }

  getBookingByID(bookingID: number): Observable<Booking> {
    const url = `${this.apiUrl}/${bookingID}`;
    return this.http.get<Booking>(url);
  }

  // Add any additional methods as per your requirements

}
