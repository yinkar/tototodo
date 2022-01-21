import { Given, Then } from 'cypress-cucumber-preprocessor/steps';

When(`i type {string} to input`, s => {
    cy.visit('/');
    cy.get('.todos').contains(s);
});

Then(`input should contains {string}`, s => {

});

Given(`i press the button`, () => {
    cy.visit('/');
});

Then(`{string} should be in the list`, s => {
    cy.get('.todos').contains(s);
});