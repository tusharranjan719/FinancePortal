import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

declare const $: any;
declare interface RouteInfo {
    path: string;
    title: string;
    icon: string;
    class: string;
}
export const ROUTES: RouteInfo[] = [
    { path: '/dashboard', title: 'Dashboard',  icon: 'dashboard', class: '' },
    { path: '/dashboard/user-profile', title: 'User Profile',  icon:'person', class: '' },
    { path: '/dashboard/table-list', title: 'Transactions',  icon:'content_paste', class: '' },
    { path: '/dashboard/notifications', title: 'Notifications',  icon:'notifications', class: '' },
    { path: '/dashboard/upgrade', title: 'Log Out',  icon:'unarchive', class: 'active-pro' },
];

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {
  menuItems: any[];

  constructor(private _router: Router) { }

  ngOnInit() {
    this.menuItems = ROUTES.filter(menuItem => menuItem);
  }
  goToItem(route: any) {
    switch(route) {
      case '/dashboard/upgrade':
        this._router.navigate(['login']);
        break;
    }
    //this._router.navigate(['dashboard/notifications'])
  }
  isMobileMenu() {
      if ($(window).width() > 991) {
          return false;
      }
      return true;
  };
}
