Feature: add todo
    Scenario: typing on input
        When i type "play guitar" to input
        Then input should contains "play guitar"

    Scenario: clicking button
        When i press the button
        Then "play guitar" should be in the list