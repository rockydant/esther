import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from 'src/app/auth.guard';

const routes: Routes = [RouterModule.forChild([
  { path: 'users', loadChildren: () => import('./user/user.module').then(m => m.UserModule), canActivate: [AuthGuard] },
])];

@NgModule({
  imports: [RouterModule.forChild([
    { path: 'users', loadChildren: () => import('./user/user.module').then(m => m.UserModule), canActivate: [AuthGuard] },  
  ])],
  exports: [RouterModule]
})
export class SettingRoutingModule { }
