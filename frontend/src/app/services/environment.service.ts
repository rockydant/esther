import { HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root',
})
export class EnvironmentService {
    constructor() {}

    public BackendEndpoint = import.meta.env.NG_APP_BACKEND_ENDPOINT;

    // roles endpoints
    public BackendRolesEndpoint = this.BackendEndpoint + 'roles';
    public BackendUsersEndpoint = this.BackendEndpoint + 'users';

    public HttpOptions = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
    };
}
