# esther

This is a simple template used to create a new project with Go (Gin) for backend and Angular for frontend.

## How to use it?

cd backend && go mod tidy && cd .. && cd frontend && npm install && ng serve

## What's inside?

- Backend: Gin + Postgres
- Frontend: Angular 17 + Primeng

### Backend

- [x] User authentication with JWT
- [x] CRUD for users
- [x] CRUD for roles
- [] CRUD for permissions
- [] Dockerfile
- [] Makefile

### Frontend

- [x] Login page
- [x] Dashboard page
- [] User list page
- [] Role list page
- [] Permission list page
- [] User detail page
- [] Role detail page
- [] Permission detail page
- [] Dockerfile
- [] Makefile

### TODO

- [ ] Add tests for backend
- [ ] Add tests for frontend
- [ ] Improve the code quality of backend
- [ ] Improve the code quality of frontend
- [ ] Add a docker-compose file to run both backend and frontend in one command
- [ ] Add a Makefile to run both backend and frontend in one command

### Contributing

- Fork it!
- Create your feature branch: `git checkout -b my-new-feature`
- Commit your changes: `git commit -am 'Add some feature'`
- Push to the branch: `git push origin my-new-feature`
- Submit a pull request :D

### History

- 0.1.0 Initial version
- 0.2.0 Added user authentication with JWT
- 0.3.0 Added CRUD for users, roles and permissions
- 0.4.0 Added Dockerfile and Makefile to backend

### Credits

- [Angular](https://angular.io/)
- [Gin](https://gin-gonic.com/)
- [Postgres](https://www.postgresql.org/)
- [Primeng](https://primefaces.org/primeng/)

### License
