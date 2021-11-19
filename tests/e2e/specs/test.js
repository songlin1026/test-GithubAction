// https://docs.cypress.io/api/introduction/api.html

describe('My First Test', () => {
  it('Visits the app root url', () => {
    const host = Cypress.env('host') ? Cypress.env('host') : 'http://127.0.0.1:3000/';
    cy.visit(host);
    cy.contains('Welcome to Your Vue.js App');
  });
});
