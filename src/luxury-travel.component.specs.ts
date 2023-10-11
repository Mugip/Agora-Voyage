import { ComponentFixture, TestBed } from '@angular/core/testing';
import { LuxuryTravelComponent } from './luxury-travel.component';

describe('LuxuryTravelComponent', () => {
  let component: LuxuryTravelComponent;
  let fixture: ComponentFixture<LuxuryTravelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ LuxuryTravelComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(LuxuryTravelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create the app', () => {
    expect(component).toBeTruthy();
  });

  // Add more test cases as needed
});
