import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { ApiService } from 'path/to/api.service';
import { TravelService } from 'path/to/travel.service';
import { TestBed } from '@angular/core/testing';
import { TestComponent } from './test.component';

describe('TestComponent', () => {
  let component: TestComponent;
  let fixture: ComponentFixture<TestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      declarations: [TestComponent],
      providers: [ApiService, TravelService]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create the test component', () => {
    expect(component).toBeTruthy();
  });

  it('should call the travel service and retrieve the travel data', () => {
    const travelService = TestBed.inject(TravelService);
    const httpMock = TestBed.inject(HttpTestingController);

    const expectedTravelData = [
      {
        destination: 'Paris',
        duration: 7,
        price: 1999
      },
      {
        destination: 'Rome',
        duration: 5,
        price: 1499
      }
    ];

    component.ngOnInit();

    const req = httpMock.expectOne('/api/travel');
    expect(req.request.method).toBe('GET');
    req.flush(expectedTravelData);

    expect(component.travelData).toEqual(expectedTravelData);

    httpMock.verify();
  });
})
