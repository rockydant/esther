import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserComponent } from './user.component';

const routes: Routes = [
  
];

@NgModule({
  imports: [RouterModule.forChild([
    { path: '', component: UserComponent }
  ])],
  exports: [RouterModule]
})
export class UserRoutingModule { }
