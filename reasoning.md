## Approach

I started by setting up a basic server first.
Before adding database or other features, I wanted to make sure
the application starts properly and logging is in place.
I added a simple health endpoint so I can quickly check
that the server is running as expected.

I added basic config loading to avoid hardcoding values.
For now, only the server port is configurable.

The server is healthy so I added a simple request logger.
This helps me see incoming requests and responses while developing,
and makes debugging easier before adding database logic.

I added database related configuration early so connection details
are not hardcoded and can be changed easily later.

I set up PostgreSQL using Docker so I can run the database locally
while working on the backend.

I added a migration to create the users table so the database structure
is defined before writing any queries.

I added sqlc configuration so database queries can be generated
in a type-safe way when repositories are implemented.
And later I added a basic query to verify that sqlc generation works
before adding actual repository logic.

Added basic user queries and regenerated sqlc code
to make sure the database layer is ready before using it.

I added a central database connection file in the repository layer
so sqlc queries can share a single connection pool.
And I verified the database connection by pinging Postgres during startup
to make sure the app can connect before serving requests.