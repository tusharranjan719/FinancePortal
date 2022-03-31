import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
@Component({
  selector: 'financex-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  emailRegx = /^(([^<>+()\[\]\\.,;:\s@"-#$%&=]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,3}))$/;
  invalidLoginCreds: boolean = false;
  invalidLoginMsg: string = '';
  constructor(private fb: FormBuilder, private http: HttpClient, private router: Router) { 
    let bodyElem: HTMLElement | null = document.getElementById('financeX');
    bodyElem!['className'] += 'login_bg';

  }

  loginForm = this.fb.group({
    email: ['', [Validators.required, Validators.pattern(this.emailRegx)]],
    password: ['', Validators.required],
  });

  signupForm = this.fb.group({
    first_name: ['', Validators.required],
    last_name: ['', Validators.required],
    email: ['', [Validators.required, Validators.pattern(this.emailRegx)]],
    password: ['', Validators.required],
    retype_password: ['', [Validators.required, this.passwordMatcher.bind(this)]]
  });

  private passwordMatcher(control: FormControl): { [s: string]: boolean } | null {
    if (
        this.signupForm &&
        (control.value !== this.signupForm.controls['password'].value)
    ) {
        return { passwordNotMatch: true };
    }
    return null;
  } 

  loginSubmit(){
    let formValue = this.loginForm.value;
    let postData = {
      'username': formValue.email,
      'password': formValue.password
    };
    this.http.post('/signIn', postData).subscribe((data)=>{
      this.invalidLoginCreds = false;
      console.log(data);
    },
    (error)=>{
      if(error.status == 401){
        this.invalidLoginCreds = true;
        this.invalidLoginMsg = error.error;
        console.log(error);
      }
      else{
        this.router.navigate(['dashboard']);
      }
      
    });
  }
  signUpSubmit(){
    let formValue = this.signupForm.value;
    //console.log(formValue);
    let postData = {
      'first_name': formValue.first_name,
      'last_name': formValue.last_name,
      'username': formValue.email,
      'password': formValue.password
    };
    this.http.post('/signUp', postData).subscribe((data)=>{
      console.log(data);
    });
  }

  get signUpFormControls() {
    return this.signupForm.controls;
  }

  get loginFormControls() {
    return this.loginForm.controls;
  }

  signUpFormHasError(c: any, err: string){
    if(this.signupForm.get(c)?.hasError(err)){
      return true;
    }
    return false;
  }

  loginFormHasError(c: any, err: string){
    if(this.loginForm.get(c)?.hasError(err)){
      return true;
    }
    return false;
  }

  ngOnInit(): void {
  }

  ngOnDestroy(){
    let bodyElem: HTMLElement | null = document.getElementById('financeX');
    bodyElem!['classList'].remove('login_bg');
  }
}