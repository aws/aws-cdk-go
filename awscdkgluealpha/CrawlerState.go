package awscdkgluealpha


// Represents the state of a crawler for a condition in the Glue Trigger API.
// Experimental.
type CrawlerState string

const (
	// The crawler is currently running.
	// Experimental.
	CrawlerState_RUNNING CrawlerState = "RUNNING"
	// The crawler is in the process of being cancelled.
	// Experimental.
	CrawlerState_CANCELLING CrawlerState = "CANCELLING"
	// The crawler has been cancelled.
	// Experimental.
	CrawlerState_CANCELLED CrawlerState = "CANCELLED"
	// The crawler has completed its operation successfully.
	// Experimental.
	CrawlerState_SUCCEEDED CrawlerState = "SUCCEEDED"
	// The crawler has failed to complete its operation.
	// Experimental.
	CrawlerState_FAILED CrawlerState = "FAILED"
	// The crawler encountered an error during its operation.
	// Experimental.
	CrawlerState_ERROR CrawlerState = "ERROR"
)

