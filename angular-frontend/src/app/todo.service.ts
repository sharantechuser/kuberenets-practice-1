import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { AppSettings, ToDo } from './home/home.component';

@Injectable({
  providedIn: 'root'
})
export class TodoService {

  constructor(private http: HttpClient) { }

  getTodoList() {

    return this.http.get<ToDo[]>(AppSettings.TODO_GET_API_ENDPOINT)
  
  }

}
