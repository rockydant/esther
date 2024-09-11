import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { jwtDecode } from 'jwt-decode';

@Injectable({
    providedIn: 'root',
})
export class AuthService {
    private tokenKey = 'token';
    private _backendEndpoint = import.meta.env.NG_APP_BACKEND_ENDPOINT;

    constructor(private http: HttpClient, private router: Router) { }

    login(username: string, password: string) {
        return this.http.post<any>(this._backendEndpoint + 'login', {
            username,
            password,
        });
    }

    logout() {
        localStorage.removeItem(this.tokenKey);
        this.router.navigate(['/login']);
    }

    isAuthenticated(): boolean {
        const token = localStorage.getItem(this.tokenKey);
        if (token) {
            // Decode JWT token to check if it’s still valid
            const decodedToken: any = jwtDecode(token);
            const expirationDate = new Date(0);
            expirationDate.setUTCSeconds(decodedToken.exp);
            return expirationDate > new Date();
        }
        return false;
    }

    getToken() {
        return localStorage.getItem(this.tokenKey);
    }
}
