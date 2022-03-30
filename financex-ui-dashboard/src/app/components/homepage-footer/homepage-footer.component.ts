import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-homepage-footer',
  templateUrl: './homepage-footer.component.html',
  styleUrls: ['./homepage-footer.component.css']
})
export class HomePageFooterComponent implements OnInit {
  test : Date = new Date();
  
  constructor() { }

  ngOnInit() {
  }

}
