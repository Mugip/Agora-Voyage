import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-contact-form',
  templateUrl: './contact-form.component.html',
  styleUrls: ['./contact-form.component.css']
})
export class ContactFormComponent {
  contact = {
    name: '',
    email: '',
    message: ''
  };

  constructor(private http: HttpClient) {}

  submitForm() {
    this.http.post('/api/contact', this.contact).subscribe(
      () => {
        alert('Thank you for contacting us!');
      },
      (error) => {
        console.error(error);
        alert('Failed to submit the contact form. Please try again later.');
      }
    );
  }
