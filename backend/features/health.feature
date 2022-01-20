Feature: check health
    Scenario: POST method not allowed
        When client makes a "POST" request to "/health"
        Then response code should be 405

    Scenario: should get an json response
        When client makes a "GET" request to "/health"
        Then response code should be 200
        And response body should match with:
            """
            {
                "alive": true
            }
            """