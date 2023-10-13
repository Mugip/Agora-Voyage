import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import {HTTP_INTERCEPTORS} from '@angular/common/http' ;

// Import components
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { NavbarComponent } from './shared/navbar/navbar.component';
import { BookingComponent } from './booking/booking.component';
import { DestinationComponent } from './destination/destination.component';

// Import services
import { AuthService } from './core/auth/auth.service';
import { BookingService } from './core/booking/booking.service';
import { DestinationService } from './core/destination/destination.service';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    SignupComponent,
    NavbarComponent,
    BookingComponent,
    DestinationComponent,
    ProtectedComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    RouterModule.forRoot([
      { path: 'home', component: HomeComponent },
      { path: 'login', component: LoginComponent },
      { path: 'signup', component: SignupComponent },
      { path: 'booking', component: BookingComponent },
      { path: 'destination', component: DestinationComponent },
      { path: '**', redirectTo: 'home', pathMatch: 'full' },
    ]),
  ],
  providers: [
    AuthService,
    BookingService,
    DestinationService,
    AuthInterceptor,
  ],
  bootstrap: [AppComponent],
})
export class AppModule { }

