describe('It should have FinanceX title', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has the correct title', () => {
    cy.title().should('equal', 'FinanceX');
  });
  it('has anchor tags', () => {
    cy.get('.scroll').should('have.length', 5)
  });
  it('has active home link', () => {
    cy.get('.navbar li:first').should('have.class', 'active')
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


