package awscodecommit


// Repository events that will cause the trigger to run actions in another service.
// Experimental.
type RepositoryEventTrigger string

const (
	// Experimental.
	RepositoryEventTrigger_ALL RepositoryEventTrigger = "ALL"
	// Experimental.
	RepositoryEventTrigger_UPDATE_REF RepositoryEventTrigger = "UPDATE_REF"
	// Experimental.
	RepositoryEventTrigger_CREATE_REF RepositoryEventTrigger = "CREATE_REF"
	// Experimental.
	RepositoryEventTrigger_DELETE_REF RepositoryEventTrigger = "DELETE_REF"
)

