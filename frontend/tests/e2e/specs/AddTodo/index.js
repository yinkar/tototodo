import { When, Then } from 'cypress-cucumber-preprocessor/steps';

When(`i type {string} to input`, s => {
    cy.visit('/');
    cy.get('#content').type(s);
});

Then(`input should contains {string}`, s => {
    cy.get('#content').should('have.value', s);
});

When(`i press the button`, () => {
    cy.get('#add').click();
});

Then(`{string} should be in the list`, s => {
    cy.get('.todos').contains(s);
});