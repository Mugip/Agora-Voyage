import { Component, OnInit } from '@angular/core';
import { DestinationService } from '../services/destination.service';

@Component({
  selector: 'app-destination',
  templateUrl: './destination.component.html',
  styleUrls: ['./destination.component.css']
})
export class DestinationComponent implements OnInit {
  destinations: any[] = [];

  constructor(private destinationService: DestinationService) {}

  ngOnInit(): void {
    this.getDestinations();
  }

  getDestinations(): void {
    this.destinationService.getDestinations().subscribe(
      (response) => {
        this.destinations = response;
      },
      (error) => {
        console.log('An error occurred while fetching destinations:', error);
      }
    );
  }

  addDestination(destination: any): void {
    this.destinationService.addDestination(destination).subscribe(
      (response) => {
        // Handle the response (if required)
        console.log('Destination added successfully');
        this.getDestinations(); // Refresh destinations after adding
      },
      (error) => {
        console.log('An error occurred while adding a destination:', error);
      }
    );
  }
}
