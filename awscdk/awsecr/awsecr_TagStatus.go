package awsecr


// Select images based on tags.
// Experimental.
type TagStatus string

const (
	// Rule applies to all images.
	// Experimental.
	TagStatus_ANY TagStatus = "ANY"
	// Rule applies to tagged images.
	// Experimental.
	TagStatus_TAGGED TagStatus = "TAGGED"
	// Rule applies to untagged images.
	// Experimental.
	TagStatus_UNTAGGED TagStatus = "UNTAGGED"
)

