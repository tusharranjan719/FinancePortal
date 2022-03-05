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



