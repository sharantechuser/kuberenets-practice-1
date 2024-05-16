import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { TodoComponent } from './todo/todo.component';

export const routes: Routes = [
    {path:'', component: HomeComponent},
    {path:'home', component: HomeComponent},
    {path: 'login', component: LoginComponent},
    {path: 'todo', component: TodoComponent},
    // { path: '**', redirectTo: 'home' } ,
];
