import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HomePageFooterComponent } from './homepage-footer.component';

describe('FooterComponent', () => {
  let component: HomePageFooterComponent;
  let fixture: ComponentFixture<HomePageFooterComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HomePageFooterComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HomePageFooterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
