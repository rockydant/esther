import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, of, tap } from 'rxjs';
import { RoleDTO } from '../models/role';
import { MessageService } from './message/message.service';
import { EnvironmentService } from './environment.service';
import { UserDTO } from '../models/user';

@Injectable({
    providedIn: 'root',
})
export class UserService {
    constructor(
        private http: HttpClient,
        private messageService: MessageService,
        private envService: EnvironmentService
    ) {}

    // Generated CRUD services for user routes in backend
    getRoles(): Observable<RoleDTO[]> {
        return this.http
            .get<RoleDTO[]>(`${this.envService.BackendRolesEndpoint}`)
            .pipe(
                tap((_) => this.log('fetched role records')),
                catchError(this.handleError<RoleDTO[]>('getRoles', []))
            );
    }

    getUsers(): Observable<UserDTO[]> {
        return this.http
            .get<UserDTO[]>(`${this.envService.BackendUsersEndpoint}`)
            .pipe(
                tap((_) => this.log('fetched user records')),
                catchError(this.handleError<UserDTO[]>('getUsers', []))
            );
    }

    createRole(name: string): Observable<RoleDTO> {
        return this.http
            .post<RoleDTO>(
                `${this.envService.BackendRolesEndpoint}/${name}`,
                this.envService.HttpOptions
            )
            .pipe(
                tap((repsonse: any) =>
                    this.log(
                        `added createRole w/ id=${name}: response ${repsonse}`
                    )
                ),
                catchError(this.handleError<RoleDTO>('createRole'))
            );
    }

    getRole(id: string): Observable<RoleDTO> {
        return this.http
            .get<RoleDTO>(`${this.envService.BackendRolesEndpoint}/${id}`)
            .pipe(
                tap((_) => this.log('fetched getRole records')),
                catchError(this.handleError<RoleDTO>('getRole', {}))
            );
    }

    updateRole(role: RoleDTO): Observable<RoleDTO> {
        return this.http
            .put<RoleDTO>(
                this.envService.BackendRolesEndpoint,
                role,
                this.envService.HttpOptions
            )
            .pipe(
                tap((role: RoleDTO) =>
                    this.log(`added updateRole w/ id=${role.Name}`)
                ),
                catchError(this.handleError<RoleDTO>('updateRole'))
            );
    }

    removeRole(id: string): Observable<any> {
        return this.http
            .delete<any>(
                `${this.envService.BackendRolesEndpoint}/${id}`,
                this.envService.HttpOptions
            )
            .pipe(
                tap((_) => this.log(`deleted AppSetting id=${id}`)),
                catchError(this.handleError<any>('deleteSettings'))
            );
    }

    private log(message: string) {
        this.messageService.add(`HeroService: ${message}`);
    }

    private handleError<T>(operation = 'operation', result?: T) {
        return (error: any): Observable<T> => {
            // TODO: send the error to remote logging infrastructure
            console.error(error); // log to console instead

            // TODO: better job of transforming error for user consumption
            this.log(`${operation} failed: ${error.message}`);

            // Let the app keep running by returning an empty result.
            return of(result as T);
        };
    }
}
