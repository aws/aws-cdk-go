package awscodecommit


// List of event types for AWS CodeCommit.
// See: https://docs.aws.amazon.com/dtconsole/latest/userguide/concepts.html#events-ref-repositories
//
// Experimental.
type RepositoryNotificationEvents string

const (
	// Trigger notication when comment made on commit.
	// Experimental.
	RepositoryNotificationEvents_COMMIT_COMMENT RepositoryNotificationEvents = "COMMIT_COMMENT"
	// Trigger notification when comment made on pull request.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_COMMENT RepositoryNotificationEvents = "PULL_REQUEST_COMMENT"
	// Trigger notification when approval status changed.
	// Experimental.
	RepositoryNotificationEvents_APPROVAL_STATUS_CHANGED RepositoryNotificationEvents = "APPROVAL_STATUS_CHANGED"
	// Trigger notifications when approval rule is overridden.
	// Experimental.
	RepositoryNotificationEvents_APPROVAL_RULE_OVERRIDDEN RepositoryNotificationEvents = "APPROVAL_RULE_OVERRIDDEN"
	// Trigger notification when pull request created.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_CREATED RepositoryNotificationEvents = "PULL_REQUEST_CREATED"
	// Trigger notification when pull request source updated.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_SOURCE_UPDATED RepositoryNotificationEvents = "PULL_REQUEST_SOURCE_UPDATED"
	// Trigger notification when pull request status is changed.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_STATUS_CHANGED RepositoryNotificationEvents = "PULL_REQUEST_STATUS_CHANGED"
	// Trigger notification when pull requset is merged.
	// Experimental.
	RepositoryNotificationEvents_PULL_REQUEST_MERGED RepositoryNotificationEvents = "PULL_REQUEST_MERGED"
	// Trigger notification when a branch or tag is created.
	// Experimental.
	RepositoryNotificationEvents_BRANCH_OR_TAG_CREATED RepositoryNotificationEvents = "BRANCH_OR_TAG_CREATED"
	// Trigger notification when a branch or tag is deleted.
	// Experimental.
	RepositoryNotificationEvents_BRANCH_OR_TAG_DELETED RepositoryNotificationEvents = "BRANCH_OR_TAG_DELETED"
	// Trigger notification when a branch or tag is updated.
	// Experimental.
	RepositoryNotificationEvents_BRANCH_OR_TAG_UPDATED RepositoryNotificationEvents = "BRANCH_OR_TAG_UPDATED"
)

