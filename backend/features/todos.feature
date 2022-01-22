Feature: todos
    Scenario: should get empty todolist
        When client makes a "GET" request to "/todos"
        Then response code should be 200
        And response body should look like:
            """
            []
            """

    Scenario: should get todos
        Given there are todos:
            | content    |
            | Cook food  |
            | Eat food   |
        When client makes a "GET" request to "/todos"
        Then response code should be 200
        And response body should look like:
            """
            [
                {
                    "Content": "Cook food"
                },
                {
                    "Content": "Eat food"
                }
            ]
            """
