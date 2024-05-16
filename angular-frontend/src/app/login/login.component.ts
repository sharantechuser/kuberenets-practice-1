import { Component } from '@angular/core';
import { Router, RouterLink, RouterOutlet } from '@angular/router';
import { MatIconModule } from '@angular/material/icon';
import { FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [RouterLink, RouterOutlet, ReactiveFormsModule, MatIconModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})

export class LoginComponent {

  // items:Object
  constructor(
    private formBuilder: FormBuilder,
    private router: Router,
    private http: HttpClient

  ) { }


  checkoutForm = this.formBuilder.group({
    username: '',
    password: ''
  });

  onSubmit(): void {
    // Process checkout data here


    console.info('Your todo added', this.checkoutForm.value);
    let userName = this.checkoutForm.value.username
    let password = this.checkoutForm.value.password
    if (password == '' || userName == '') {
      console.log("Missing Username or password");
      return
    }
    
    this.checkoutForm.reset();

    this.http.get(AppSettings.LOGIN_API_ENDPOINT)
      .subscribe((data) => {
        console.log(data); // Process your data here
        // TODO if login success
        let user = {'username': `${userName}`, 'loggedin': 'true'}
        localStorage.setItem(AppSettings.LOCAL_STORAGE_USER_KEY, user.username )
        localStorage.setItem(AppSettings.LOCAL_STORAGE_LOGGEDIN_KEY, user.loggedin )
        this.router.navigate(['/home']);
      });


  }
}

export class AppSettings {
  public static LOGIN_API_ENDPOINT = 'https://dummyjson.com/products';
  public static LOCAL_STORAGE_USER_KEY = 'current_user';
  public static LOCAL_STORAGE_LOGGEDIN_KEY = 'loggedIn';
}