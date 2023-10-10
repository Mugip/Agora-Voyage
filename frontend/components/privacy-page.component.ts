typescript
// privacy-page.component.ts
import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-privacy-page',
  templateUrl: './privacy-page.component.html',
  styleUrls: ['./privacy-page.component.css']
})
export class PrivacyPageComponent implements OnInit {
  privacyContent: string;

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.fetchPrivacyContent();
  }

  fetchPrivacyContent() {
    this.http.get('/api/privacy', { responseType: 'text' }).subscribe(
      (content) => {
        this.privacyContent = content;
      },
      (error) => {
        console.error(error);
        // Handle error
      }
    );
  }
