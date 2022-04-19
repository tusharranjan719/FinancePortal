describe('It should have FinanceX title', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has the correct title', () => {
    cy.title().should('equal', 'FinanceX');
  });
});
describe('It should have 5 nav links', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has anchor tags', () => {
    cy.get('.scroll').should('have.length', 5)
  });
  it('has active home link', () => {
    cy.get('.navbar li:first').should('have.class', 'active')
  });
});
describe('It should have nav texts', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has Home text in nav', () => {
    cy.get('.scroll.active > a').should('contain', 'Home')
  });
  it('has About text in nav', () => {
    cy.get('.navbar li:nth-child(n+2):nth-child(-n+2) > a').should('contain', 'About')
  });
  it('has Features text in nav', () => {
    cy.get('.navbar li:nth-child(n+3):nth-child(-n+3) > a').should('contain', 'Features')
  });
  it('has Team text in nav', () => {
    cy.get('.navbar li:nth-child(n+4):nth-child(-n+4) > a').should('contain', 'Team')
  });
  it('has Login text in nav', () => {
    cy.get('.navbar li:nth-child(n+5):nth-child(-n+5) > a').should('contain', 'Login')
  });
});
describe('It should load logo image', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has Home text in nav', () => {
    cy.get('[alt="logo"]').should('be.visible')
  });
});
describe('It should have background image', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has background-image', () => {
    cy.get('[id^=hero-banner]').should('have.css', 'background-image')
      .and('include', 'banner.jpg')
  });
});
describe('It should have One Stop container text', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has One Stop center text', () => {
    cy.get('.banner-inner h2 > b').should('contain', 'One Stop')
  });
});
describe('It should have Finance Solution container text', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has Finance Solution center text', () => {
    cy.get('.banner-inner h2').should('contain', ' Finance Solution')
  });
});
describe('It should have Money Management Made Easy container text', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has Sign Up button', () => {
    cy.get('.banner-inner p').should('contain', 'Money Management Made Easy')
  });
});
describe('It should have Sign Up link', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has Sign Up button', () => {
    cy.get('.banner-inner a').should('contain', 'Sign Up')
  });
});
describe('It should have about container', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has about container', () => {
    cy.get('[id^=about]').should('be.visible');
  });
  it('has about image', () => {
    cy.get('[id^=about] img.img-responsive').should('be.visible');
  });
});
describe('It should have features container', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has features container', () => {
    cy.get('[id^=features]').should('be.visible');
  });
});
describe('It should show all 3 features', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has 3 features', () => {
    cy.get('[id^=features] .features').children().should('have.length', 3);
  });
  it('first feature has shared expenses', () => {
    cy.get('[id^=features] .features .media.service-box').eq(0).should('contain', 'Shared Expenses');
  });
});
describe('It should show all 3 features', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has 3 features', () => {
    cy.get('[id^=features] .features').children().should('have.length', 3);
  });
  it('first feature has shared expenses', () => {
    cy.get('[id^=features] .features .media.service-box').eq(0).should('contain', 'Shared Expenses');
  });
});
describe('Dashboard page should have 5 li nav items', () => {
  beforeEach(() => {
    cy.visit('/#/dashboard/dashboard');
  });
  it('has 5 sidebar navs', () => {
    cy.get('.sidebar-wrapper .nav').find('li').should('have.length', 5);
  });
});

describe('Dashboard page should have dashboard link active', () => {
  beforeEach(() => {
    cy.visit('/#/dashboard/dashboard');
  });
  it('dashboard link active', () => {
    cy.get('.sidebar-wrapper .nav li:first').should('have.class', 'active');
  });
});

// Functional Test Cases

describe('login url', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('redirects to another page on click', () => {
    // cy.on('url:changed', (url) => {
    //   expect(url).to.contain("login")
    // });
    cy.get('.scroll [id^=fin-login]').click({ force: true });
    //cy.wait(50);
    cy.get('.mat-card-header').should('be.visible');
  })

  describe('login url', () => {
    beforeEach(() => {
      cy.visit('/#/login');
    });
    it('first feature has shared expenses', () => {
      cy.get('.mat-card-header').should('be.visible');
    });

  });

});

describe('login form submit redirect to user dashboard', () => {
  beforeEach(() => {
    cy.visit('/#/login');
  });
  it('submit valid form', () => {
    cy.on('url:changed', (url) => {
      expect(url).to.contain("dashboard")
    });
    cy.get('.mat-form-field-infix [formcontrolname^=email]').type('user@gg.com');
    cy.get('.mat-form-field-infix [formcontrolname^=password]').type('user@gg.com');
    cy.wait(100);
    cy.get('form').submit();
  });
});

describe('user profile link click redirects to its page', () => {
  beforeEach(() => {
    cy.visit('/#/dashboard/dashboard');
  });
  it('redirects user profile', () => {
      cy.on('url:changed', (url) => {
          expect(url).to.contain("user-profile")
      });
    cy.get('.sidebar-wrapper .nav li:nth-child(n+2):nth-child(-n+2) a').click({ force: true });
  })
});

describe('user profile link click redirects to its page', () => {
  beforeEach(() => {
    cy.visit('/#/dashboard/dashboard');
  });
  it('redirects user profile', () => {
      cy.on('url:changed', (url) => {
          expect(url).to.contain("table-list")
      });
    cy.get('.sidebar-wrapper .nav li:nth-child(n+3):nth-child(-n+3) a').click({ force: true });
  })
});

describe('user profile link click redirects to its page', () => {
  beforeEach(() => {
    cy.visit('/#/dashboard/dashboard');
  });
  it('redirects user profile', () => {
      cy.on('url:changed', (url) => {
          expect(url).to.contain("notifications")
      });
    cy.get('.sidebar-wrapper .nav li:nth-child(n+4):nth-child(-n+4) a').click({ force: true });
  })
});

describe('sign In', () => {
  it('sign In', () => {
    cy.intercept('POST', '/signIn', {
      statusCode: 201
    }).as('new-user')
    cy.request('POST', '/signIn', {
      name: 'Sankalp',
      password: 'tdlr',
    }).then((response: any) => {
      cy.log(response)
      console.log(response)
      //expect(response.statusCode).to.eq(201)
      expect(response.body['name']).to.be.eq('Sankalp');
    })
    cy.request('POST', '/new-user', {
      name: 'Sankalp',
      password: 'tdlr',
    }).then((response: any) => {
      cy.log(response)
      console.log(response)
    })
    cy.get('@new-user').then(console.log)
  })

})

describe('sign Up New User', () => {
  it('sign Up', () => {
    cy.wait(2000)
    cy.intercept('POST', '/signUp', {
      statusCode: 201
    }).as('new-user')
    cy.request('POST', '/signUp', {
      name: 'Sankalp',
      password: 'tdlr',
    }).then((response: any) => {
      cy.log(response)
      console.log(response)
      //expect(response.statusCode).to.eq(201)
      expect(response.body['name']).to.be.eq('Sankalp');
    })
    cy.request('POST', '/new-user', {
      name: 'Sankalp',
      password: 'tdlr',
    }).then((response: any) => {
      cy.log(response)
      console.log(response)
    })
    cy.get('@new-user').then(console.log)
  })

})


describe('get all users info', () => {
  beforeEach(() => {
    cy.visit('/#/dashboard/dashboard');
  });
  it('get all user', () => {
    cy.wait(2000)
    cy.request('GET', '/users').then((response) => {
      expect(response.body).to.be.a('array');
    })
  });

});

describe('get user by id', () => {
  beforeEach(() => {
    cy.visit('/#/dashboard/dashboard');
  });
  it('get user', () => {
    cy.request('GET', '/users/1').then((response) => {
      expect(response.body['name']).to.be.eq('Sankalp Pandey');
    })
  });

});

describe('get user billing info by id', () => {
  beforeEach(() => {
    cy.visit('/#/dashboard/dashboard');
  });
  
  it('get user billing info by id', () => {
      cy.wait(2000);
      cy.request('GET', '/user_billing_info/1').then((response) => {
          expect(response.body['balance']).to.be.eq(5000);
      })
  });
});

describe('new bill split group', () => {
  it('new bill split group', () => {
    cy.wait(2000)
    cy.intercept('POST', '/billsplit/new', {
      statusCode: 201
    }).as('new-bill')
    cy.request('POST', '/billsplit/new', {
      name: 'billsplitgroup1',
    }).then((response: any) => {
      cy.log(response)
      console.log(response)
      //expect(response.statusCode).to.eq(201)
      expect(response.body['name']).to.be.eq('billsplitgroup1');
    })
    cy.get('@new-bill').then(console.log)
  })
});

describe('create new expense', () => {
  it('create new expense', () => {
    cy.wait(2000)
    cy.intercept('POST', '/billsplit/1/expense/new', {
      statusCode: 201
    }).as('new-expense')
    cy.request('POST', '/billsplit/1/expense/new', {
      name: 'billsplitgroup1',
      amount: 2000,
      payerName: "Sankalp",
      billSplitId: 2
    }).then((response: any) => {
      cy.log(response)
      console.log(response)
      //expect(response.statusCode).to.eq(201)
      expect(response.body['payerName']).to.be.eq('Sankalp');
    })
    cy.get('@new-expense').then(console.log)
  })
});

describe('get all expenses by billsplitid', () => {
  beforeEach(() => {
    cy.visit('/#/dashboard/dashboard');
  });
  
  it('gets expense by billsplitid', () => {
      cy.wait(2000);
      cy.request('GET', '/billsplit/1/expenses').then((response) => {
          expect(response.body['payerName']).to.be.eq("Sankalp");
      })
  });
});





