import { Component, OnInit } from '@angular/core';
import { Router, RouterLink, RouterOutlet } from '@angular/router';
import { MatIconModule } from '@angular/material/icon';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { NgFor } from '@angular/common';
import { HttpClient } from '@angular/common/http'
import { formatDate } from '@angular/common';


import { TodoService } from '../todo.service';


@Component({
  selector: 'app-home',
  standalone: true,
  imports: [RouterLink, RouterOutlet, NgFor, ReactiveFormsModule, MatIconModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {

  constructor(
    private todoService: TodoService,
    private formBuilder: FormBuilder,
    private router: Router,
    private http: HttpClient,

  ) {

    let loggedIn = localStorage.getItem(AppSettings.LOCAL_STORAGE_LOGGEDIN_KEY);

    if (loggedIn != "true") {
      this.router.navigate(['/login']);
    }
  }

  items: ToDo[] = []
  checkoutForm: any = this.formBuilder.group
  current_user: any = ""


  ngOnInit(): void {

    this.current_user = localStorage.getItem(AppSettings.LOCAL_STORAGE_USER_KEY);

    this.checkoutForm = this.formBuilder.group({
      user: this.current_user,
      description: '',
      datetime: [formatDate(new Date(), 'yyyy-MM-dd HH:mm', 'en'), [Validators.required]]
    });


    this.loadData()
  }

  loadData(): void {
    this.todoService.getTodoList().subscribe((todos: ToDo[]) => {
      this.items = todos
    })
  }

  onSubmit(): void {
    console.log(this.checkoutForm.value);
    if (this.checkoutForm.value.datetime == null || this.checkoutForm.value.description == null || this.checkoutForm.value.description == "") {
      console.log("Missing description");
      return
    }
    let NewTime = formatDate(this.checkoutForm.value.datetime, 'yyyy-MM-dd HH:mm', 'en')
    let todo = { 'user': `${this.current_user}`, 'description': `${this.checkoutForm.value.description}`, 'datetime': `${this.checkoutForm.value.datetime}` }
    // this.items.push(todo)
    console.info('Your todo added', todo);

    this.http.post(AppSettings.TODO_POST_API_ENDPOINT, todo)
      .subscribe((response) => {
        console.log(response); // Handle the response from the server
        this.loadData()
      });

    this.checkoutForm.reset();
  }

  onLogout() {
    localStorage.clear()
    this.router.navigate(['/login']);
  }

  currentDate() {
    let time = new Date();
    let year = time.getFullYear();
    let month = time.getMonth() + 1
    let day = time.getDate();

    let NewTime = year + "-" + ("00" + month).slice(-2) + "-" + ("00" + day).slice(-2)
    return NewTime
  }
}

export class ToDo {
  user: string;
  description: string;
  datetime: string;

  constructor(user: string, description: string, datetime: string) {
    this.user = user;
    this.description = description;
    this.datetime = datetime;
  }
}


export class AppSettings {
  public static TODO_POST_API_ENDPOINT = 'http://127.0.0.1:4000/api/todo';
  public static TODO_GET_API_ENDPOINT = 'http://127.0.0.1:4000/api/todos';
  public static LOCAL_STORAGE_USER_KEY = 'current_user';
  public static LOCAL_STORAGE_LOGGEDIN_KEY = 'loggedIn';
}
