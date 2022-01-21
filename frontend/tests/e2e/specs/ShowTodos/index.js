import { Given, Then } from 'cypress-cucumber-preprocessor/steps';

Given(`client on the main page`, () => {
    cy.visit('/');
    cy.the('header').contains('Tototo Do');
    cy.get('#app div').children('.todos');
});

Then(`todos should list`, () => {
    cy.get('.todos')
});