import { Component, OnInit } from '@angular/core';
import { ApiService } from 'path/to/api.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {
  data: any;

  constructor(private apiService: ApiService) { }

  ngOnInit() {
    this.fetchData();  // Calling the fetchData method on component initialization
  }

  fetchData() {
    this.apiService.getData().subscribe(
      (response: any) => {
        this.data = response;
      },
      (error: any) => {
        console.error('Error fetching data:', error);
      }
    );
  }
