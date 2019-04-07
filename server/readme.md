# Vice Software Go API Boilerplate

This boilerplate was created to provided an opinonated starting point for projects that share our development values and priorities.

- [Vice Software Go API Boilerplate](#vice-software-go-api-boilerplate)
- [Getting Started](#getting-started)
	- [Installing Postgres](#installing-postgres)
		- [Starting the Database](#starting-the-database)
		- [Running PSQL](#running-psql)
	- [Setting Up the Database](#setting-up-the-database)
		- [Creating User](#creating-user)
		- [Creating the DB](#creating-the-db)
		- [Initializing the DB Schema](#initializing-the-db-schema)
		- [Seeding the Database](#seeding-the-database)
- [Installing Depedencies, Building and Running the App](#installing-depedencies-building-and-running-the-app)
- [Changing default configurations](#changing-default-configurations)
- [Our Values and Priorities](#our-values-and-priorities)

# Getting Started

1. Install Go
2. [Install PostGres Version 10](#installing-postgres)
3. [Setup the database](#setting-up-the-database)

## Installing Postgres

On mac we recommend using [Postico](https://eggerapps.at/postico/docs/v1.5.6/) (GUI app for working with Postgres) and [Homebrew](https://brew.sh/) to install [Postgres 10](https://formulae.brew.sh/formula/postgresql@10#default) as shown below.

```
brew install postgresql@10
```

> Note: There is a [good setup guide located here](https://medium.freecodecamp.org/how-to-get-started-with-postgresql-9d3bc1dd1b11) but make sure you that you use the above command when installing Postgres as you need version 10.

To get post gres commandline setup after installing with brew execute

```
echo 'export PATH="/usr/local/opt/postgresql@10/bin:$PATH"' >> ~/.bash_profile
```

To test if you successfully installed Postgrew **open a new terminal windows** and execute the following command

```
postgres --version
```

If you see the following then Postgres is installed!

```
postgres (PostgreSQL) 10.6
```

### Starting the Database

If you installed using brew as described above then execute the following command to start Postgres version 10.

```
brew services start postgresql@10
```

To verify that the DB started [run PSQL](#running-PSQL) as decribed below.

### Running PSQL

PSQL allows running PSQL commands. To run PSQL use the command below.

```
psql postgres
```

and you should see something like this

```
psql (10.6)
Type "help" for help.

postgres=#
```

This is the `psql` client. To exit type

```
/q
```

and hit enter.

## Setting Up the Database

After you have Postgress installed and available on the commandline issue the following commands to setup the user and database.

> Note: This guide uses default configurations in the API for DB Name, username, password, etc... see [Changing default configurations](#changing-default-configurations) below for details on how to change the defaults that are USED by the API.

### Creating User

```
createuser -P -e vice_boilerplate_user
```

Enter `vicesofware` as the password when prompted to use boilerplate defaults.

To verify the user was created execute `\du` in psql and you should see something like shown below.

```
vice_boilerplate=# \du
                                         List of roles
       Role name       |                         Attributes                         | Member of
-----------------------+------------------------------------------------------------+-----------
 ryan.vice             | Superuser, Create role, Create DB, Replication, Bypass RLS | {}
 vice_boilerplate_user |                                                            | {}
```

### Creating the DB

```
createdb -O vice_boilerplate_user vice_boilerplate
```

To verify the user was created execute `\l` in psql and you should see something like shown below.

```
vice_boilerplate=# \l
                                               List of databases
       Name       |         Owner         | Encoding |   Collate   |    Ctype    |      Access privileges
------------------+-----------------------+----------+-------------+-------------+-----------------------------
 postgres         | ryan.vice             | UTF8     | en_US.UTF-8 | en_US.UTF-8 |
 template0        | ryan.vice             | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =c/"ryan.vice"             +
                  |                       |          |             |             | "ryan.vice"=CTc/"ryan.vice"
 template1        | ryan.vice             | UTF8     | en_US.UTF-8 | en_US.UTF-8 | =c/"ryan.vice"             +
                  |                       |          |             |             | "ryan.vice"=CTc/"ryan.vice"
 vice_boilerplate | vice_boilerplate_user | UTF8     | en_US.UTF-8 | en_US.UTF-8 |
```

### Initializing the DB Schema

To get the DB Schema needed for the boilerplate setup exuecte the following SQL commands.

```SQL
CREATE TABLE contacts (
	id              SERIAL PRIMARY KEY,
	first_name           VARCHAR(100) NOT NULL,
	last_name  VARCHAR(100) NOT NULL,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone null
);

CREATE TABLE addresses (
	id SERIAL PRIMARY KEY,
	contact_id int not null,
	line1 varchar(100) not null,
	line2 varchar(100) null,
	city varchar(50) not null,
	state_province varchar(50) not null,
	postal_code varchar(50) not null,
	created_at timestamp with time zone not null,
	updated_at timestamp with time zone null
);

alter table addresses
add constraint fk_addresses_contact_id
foreign key (contact_id)
references contacts(id);
```

### Seeding the Database

To seed the database run the following queries.

> Note: `contact_id` in the `addresses` rows must match an `id` in the `contacts` table.

```SQL
INSERT INTO contacts (first_name, last_name, created_at, updated_at)
VALUES ('ryan', 'vice', current_timestamp, current_timestamp);

INSERT INTO contacts (first_name, last_name, created_at, updated_at)
VALUES ('prashanth', 'tondapu', current_timestamp, current_timestamp);

INSERT INTO addresses (
    contact_id,
    line1,
    line2,
    city,
    state_province,
    postal_code,
    created_at,
    updated_at
) VALUES (
	1, -- This must match id in contacts table
	'679  Strother Street',
	'suite 3000',
	'Calera',
	'Alabama',
	'35040',
	current_timestamp,
	current_timestamp

);

INSERT INTO addresses (
    contact_id,
    line1,
    line2,
    city,
    state_province,
    postal_code,
    created_at,
    updated_at
) VALUES (
	1,
    '4306  Penn Street',
    'apt 201',
    'Oates',
    'Missouri',
    '63625',
	current_timestamp,
	current_timestamp

);

INSERT INTO addresses (
    contact_id,
    line1,
    line2,
    city,
    state_province,
    postal_code,
    created_at,
    updated_at
) VALUES (
    2,
    '3339  Woodland Drive',
    'suite 621',
    'Spencer',
    'Iowa',
    '51301',
    current_timestamp,
    current_timestamp

);


```

# Installing Depedencies, Building and Running the App

Installing dependencies and Buidling can both be done by executing the command below

`go build`

from the `./cmd/webserver` directory.

To run the app simply execute

`./webserver`

from the `./cmd/webserver` directory.

Once you run the server you can see that it's working by opening the swagger page in the browser at http://127.0.0.1:8423/swagger/index.html

# Changing default configurations

If you want to configure a different databasename, username, password, etc... then make sure you update the database settings shown below in `./cmd/webserver/main.go` to be consistent with what you want to use.

```go
var (
	app = kingpin.New("skeleton", "A skeleton REST API that uses Postgres.")

	flagListen     = app.Flag("listen", "The HTTP listen address.").Default("127.0.0.1:8423").String()
	flagDBHost     = app.Flag("db-host", "The database host.").Default("127.0.0.1").String()
	flagDBPort     = app.Flag("db-port", "The database port.").Default("5434").Int()
	flagDBUser     = app.Flag("db-user", "The database user.").Default("postgres").String()
	flagDBPassword = app.Flag("db-password", "The database user's password.").Default("password").String()
	flagDBName     = app.Flag("db-name", "The database name.").Default("vicetestdb").String()
	flagDBSSL      = app.Flag("db-ssl", "The database SSL mode.").Default("disable").String()
)
```

# Our Values and Priorities

Software is all about tradeoffs. The boilerplate for for projects and teams who:

1. Want to get started quickly with a GO API
2. Want to use [Gorrilla Mux](https://www.gorillatoolkit.org/pkg/mux) for their outing
3. Want to use [POSTGres](https://www.postgresql.org/about/news/1786/)
4. Like an ORM that provides some higher abstractions like [GORM](http://gorm.io/)
5. Prefer [Go Modules](https://github.com/golang/go/wiki/Modules)
6. Prefer a thin architecture stack without a lot of layers. Note that you can add layers as your app grows complex but this boilerplate provides a thin starting point
