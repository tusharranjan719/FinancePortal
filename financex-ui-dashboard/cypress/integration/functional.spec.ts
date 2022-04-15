describe('login form submit', () => {
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

describe('get user', () => {
    beforeEach(() => {
        cy.visit('/#/dashboard/dashboard');
    });
    it('get user', () => {
        cy.request('GET', '/users').then((response) => {
            expect(response.body).to.be.a('array');
        })
    });
    
});
