import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private _backendEndpoint = import.meta.env.NG_APP_BACKEND_ENDPOINT;

  constructor(private http: HttpClient, private router: Router) { }
  
}
