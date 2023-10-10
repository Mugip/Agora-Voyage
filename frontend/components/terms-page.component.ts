typescript
// terms-page.component.ts
import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-terms-page',
  templateUrl: './terms-page.component.html',
  styleUrls: ['./terms-page.component.css']
})
export class TermsPageComponent implements OnInit {
  termsContent: string;

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.fetchTermsContent();
  }

  fetchTermsContent() {
    this.http.get('/api/terms', { responseType: 'text' }).subscribe(
      (content) => {
        this.termsContent = content;
      },
      (error) => {
        console.error(error);
        // Handle error
      }
    );
  }
