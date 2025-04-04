package awscodecommit


// Repository events that will cause the trigger to run actions in another service.
type RepositoryEventTrigger string

const (
	RepositoryEventTrigger_ALL RepositoryEventTrigger = "ALL"
	RepositoryEventTrigger_UPDATE_REF RepositoryEventTrigger = "UPDATE_REF"
	RepositoryEventTrigger_CREATE_REF RepositoryEventTrigger = "CREATE_REF"
	RepositoryEventTrigger_DELETE_REF RepositoryEventTrigger = "DELETE_REF"
)

