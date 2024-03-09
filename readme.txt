GeoNote is a location-based application that allows users to find frinds based on their geographical location.


User Authentication and Authorization:

1. Implement user registration and login functionality.
2. Use JWT (JSON Web Tokens) for secure authentication.
3. Ensure that users can only access and modify their own data.


GraphQL API:
Location-Based Functionality:
1. Integrate a geocoding service to convert user-provided addresses into latitude and longitude coordinates.
2. Utilize the Haversine formula to calculate distances between two sets of coordinates.
3. Allow users to set a preferred location and receive notifications or see notes when they enter a specified radius around that location.

NoSQL Database:
1. Use MongoDB for efficient storage and retrieval of location-based data.

Real-Time Updates:
1. Use WebSocket or a similar technology to push updates to users when new notes are created or existing notes are modified.

Scalability:
1.Optimize queries and database indexing for efficient retrieval of location-based data.
2. Running the application through a k8s pods...

Testing:
1. Write unit tests and integration tests for the backend functionality.
2. Test the application scenarios with BDD test.
