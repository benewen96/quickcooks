# QuickCooks / User Management

The user management subdomain provides functionality for managing users, tenants, and the assignment of user permissions

## Development

Start the database container with a postgres database:

```sh
docker-compose up -d
```

Seed the required data into the database:

```sh
go run . -seed=required
```

Run the server:

```sh
go run . -serve
```

For more information on flags run:

```sh
go run . -help
```

## Domain Model Overview

### [Models](./models/main.go)

#### [User](./models/user.go)

- A registered user within the QuickCooks domain

#### [Tenant](./models/tenant.go)

- A grouping of users to which all other application resources are associated

#### [Role](./models/role.go)

- A grouping of permissions that can be assigned to a user within a tenant

#### [Role Assignment](./models/role-assignment.go)

- An assignment of a role to a user, within a tenant

#### [Permission](./models/permission.go)

- A table-operation pair that defines an application permission

#### [Role Permission](./models/role-permission.go)

- An assignment of a permission to a role granting users assigned that role those permissions

### [Repositories](./repositories/main.go)

#### [User](./repositories/user.go)

- A registered user within the QuickCooks domain

#### [Tenant](./repositories/tenant.go)

- A grouping of users to which all other application resources are associated

#### [Role](./repositories/role.go)

- A grouping of permissions that can be assigned to a user within a tenant

#### [Role Assignment](./repositories/role-assignment.go)

- An assignment of a role to a user, within a tenant

#### [Permission](./repositories/permission.go)

- A table-operation pair that defines an application permission

#### [Role Permission](./repositories/role-permission.go)

- An assignment of a permission to a role granting users assigned that role those permissions

### [Services](./services/main.go)

#### [Registration](./services/registration.go)

- Exposes functionality to create a new QuickCooks user and to validate emails

#### [Authentication](./services/authentication.go)

- Exposes functionality to validate a user's login credentials and to validate or refresh auth tokens

#### [Authorization](./services/authorization.go)

- Exposes fucntionality to retrieve application authorization information

#### [My Profile](./services/my-profile.go)

- Exposes functionality to update a user's account information and to delete user accounts entirely

#### [My Tenants](./services/my-tenants.go)

- Exposes functionality to manage a user's tenant information and assign roles to other users
