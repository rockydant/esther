import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../../../services/auth/auth.service';
import { ButtonModule } from 'primeng/button';
import { NgIf } from '@angular/common';
import { CommonModule } from '@angular/common';
import { CheckboxModule } from 'primeng/checkbox';
import { PasswordModule } from 'primeng/password';
import { InputTextModule } from 'primeng/inputtext';

@Component({
    selector: 'app-login',
    standalone: true,
    imports: [
        NgIf,
        CommonModule,
        FormsModule,
        ButtonModule,
        CheckboxModule,
        InputTextModule,
        PasswordModule,
    ],
    templateUrl: './login.component.html',
    styleUrl: './login.component.css',
})
export class LoginComponent {
    loginData = {
        username: '',
        password: '',
    };

    errorMessage = '';

    constructor(private router: Router, private authService: AuthService) {}

    onSubmit() {
        this.authService
            .login(this.loginData.username, this.loginData.password)
            .subscribe(
                (response) => {
                    console.log(response);
                    localStorage.setItem('token', response.token);

                    // Redirect to home or dashboard after successful login
                    this.router.navigate(['/dashboard']);
                },
                (error) => {
                    // Handle error and show error message
                    this.errorMessage = 'Invalid username or password';
                }
            );
    }
}
