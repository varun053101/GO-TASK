## Approach

I started by setting up a basic server first.
Before adding database or other features, I wanted to make sure
the application starts properly and logging is in place.
I added a simple health endpoint so I can quickly check
that the server is running as expected.

I added basic config loading to avoid hardcoding values.
For now, only the server port is configurable.