create table guides(
    id         bigserial     not null,
    title      varchar(1023)  not null,
    url        varchar(1023) not null,
    primary key (id)
);

insert into guides(title, url) values
    ('Getting started with Testcontainers', 'https://testcontainers.com/getting-started/'),
    ('Getting started with Testcontainers for Java', 'https://testcontainers.com/guides/getting-started-with-testcontainers-for-java/'),
    ('Getting started with Testcontainers for .NET', 'https://testcontainers.com/guides/getting-started-with-testcontainers-for-dotnet/'),
    ('Getting started with Testcontainers for Node.js', 'https://testcontainers.com/guides/getting-started-with-testcontainers-for-nodejs/'),
    ('Getting started with Testcontainers for Go', 'https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/'),
    ('Testcontainers container lifecycle management using JUnit 5', 'https://testcontainers.com/guides/testcontainers-container-lifecycle/');