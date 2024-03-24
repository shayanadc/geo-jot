Feature: Access to the API
  In order to get access to the resources
  As a authenticated user
  I should see the json response

  Scenario: UnAuthorized User
    When I send "post" request to "/graphql" with body:
      """
      {
        "query": "{health}"
      }
      """
    Then I should recieve "401"

  Scenario: Authorized User
    Given I am an authenticated user
    When I send "post" request to "/graphql" with body:
      """
      {
        "query": "{health}"
      }
      """
    Then I should recieve "200"