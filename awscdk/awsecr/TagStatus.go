package awsecr


// Select images based on tags.
type TagStatus string

const (
	// Rule applies to all images.
	TagStatus_ANY TagStatus = "ANY"
	// Rule applies to tagged images.
	TagStatus_TAGGED TagStatus = "TAGGED"
	// Rule applies to untagged images.
	TagStatus_UNTAGGED TagStatus = "UNTAGGED"
)

