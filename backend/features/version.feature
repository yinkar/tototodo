Feature: get version
    Scenario: POST method not allowed
        When client makes a POST request to /health
        Then the response code should should be 405
        And no response should return

    Scenario: should get an json response
        When client makes a GET request to /version
        Then the response code should be 200
        And the response should match with:
            """
            {
                "version": "v0.1.0"
            }
            """