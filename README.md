# bookstore
Web development with Go
DDD approach

DDD reference:
https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1


cassandra:

```
CREATE KEYSPACE oauth WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1};
USE oauth;
CREATE TABLE access_tokens(access_token varchar PRIMARY KEY, user_id bigint, client_id bigint, expires bigint);
```
