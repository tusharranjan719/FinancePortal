import { Component, OnInit } from '@angular/core';
import { Routes } from '@angular/router';
import { FooterComponent } from '../footer/footer.component';
import { SidebarComponent } from '../sidebar/sidebar.component';


const routes: Routes = [
  { path: 'footer', component: FooterComponent },
  { path: 'sidebar', component: SidebarComponent }
];


@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
  }

}
