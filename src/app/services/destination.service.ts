import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DestinationService {
  private apiUrl = 'http://example.com/api'; // Replace with your API endpoint

  constructor(private http: HttpClient) {}

  // Example API call to get destinations
  getDestinations(): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/destinations`);
  }

  // Example API call to add a destination
  addDestination(destination: any): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/destinations`, destination);
  }
