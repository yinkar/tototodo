Feature: get version
    Scenario: POST method not allowed
        When client makes a "POST" request to "/version"
        Then response code should be 405

    Scenario: should get version number
        When client makes a "GET" request to "/version"
        Then response code should be 200
        And response body should match with:
            """
            {
                "version": "v0.1.0"
            }
            """