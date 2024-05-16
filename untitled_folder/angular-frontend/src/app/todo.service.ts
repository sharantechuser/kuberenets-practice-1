import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class TodoService {

  constructor() { }

  getTodoList() {
    return [
      { 'user': 'sharan', 'description': 'My first task', 'date': '2024-05-14 14:50' },
      { 'user': 'sharan', 'description': 'My second task', 'date': '2024-05-14 14:50' },
      { 'user': 'sharan', 'description': 'My third task', 'date': '2024-05-14 14:50' },
    ];
  }

}
