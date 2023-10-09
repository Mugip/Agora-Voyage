typescript
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Destination } from './destinationmodel';

@Injectable({
  providedIn: 'root'
})
export class DestinationService {
  private apiUrl = 'http://localhost:8000/destinations'; 
  // replace with actual API of golang server

  constructor(private http: HttpClient) { }

  getDestinations(): Observable<Destination[]> {
    return this.http.get<Destination[]>(this.apiUrl);
  }
}

// luxury-travel.service.ts
export class LuxuryTravelService {
  private unsplashApiBaseUrl = 'https://api.unsplash.com';
  private accessKey = '0se3s7GsJ_2q-8mk-BsW4H3BK5qW1FCBackbzzBcLNw';

  constructor(private httpClient: HttpClient) { }

  getRandomDestinationImage(): Observable<any> {
    const url = `${this.unsplashApiBaseUrl}/photos/random?query=destination&client_id=${this.accessKey}`;
    return this.httpClient.get<any>(url);
  }

  // Add more methods to fetch images for other sections as needed
}
