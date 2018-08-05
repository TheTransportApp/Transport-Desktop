import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { ServerManagerComponent } from './server-manager/server-manager.component';
import { TransferComponent } from './transfer/transfer.component';
import { KeysComponent } from './keys/keys.component';

import { RouterModule, Routes } from '@angular/router';
import {FormsModule} from "@angular/forms";

const appRoutes: Routes = [
  { path: 'server-manager', component: ServerManagerComponent },
  { path: 'transfer',      component: TransferComponent },
  { path: 'keys',      component: KeysComponent },
  { path: '',
    redirectTo: '/server-manager',
    pathMatch: 'full'
  },
  { path: '**', component: ServerManagerComponent }
];

@NgModule({
  declarations: [
    AppComponent,
    ServerManagerComponent,
    TransferComponent,
    KeysComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true } // <-- debugging purposes only
    )
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
