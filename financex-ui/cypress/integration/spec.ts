describe('My First Test', () => {
  beforeEach(() => {
    cy.visit('/');
  });
  it('has the correct title', () => {
    cy.title().should('equal', 'FinanceX');
  });
});


