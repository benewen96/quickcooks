// # Services
//
// The services package provides all business logic functionality to the API
// controllers
//
// Each service is responsible for a logically separated set of functionality
// that reflects a business requirements not data requirements
//
// A service should only have access to the database via repositories required
// to operate and should only have access to the necessary repositories
//
// The services within this subdomain are:
//   - Registration		(Provides functionality for creating/removing a QuickCooks account)
//   - Authentication	(Provides functionality for validating a user's credentials)
//   - Authorization	(Provides functionality for validating a user's access to resources)
//   - My Profile		(Provides functionality for managing a user's account information)
//   - My Tenants		(Provides functionality for managing a user's tenant's information)
package services
