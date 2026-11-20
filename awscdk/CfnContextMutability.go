package awscdk


// Change-safety level for a resource or resource property.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-metadata.html#aws-attribute-metadata-context-schema
//
type CfnContextMutability string

const (
	// The value must never change.
	CfnContextMutability_MUST_NEVER_CHANGE CfnContextMutability = "MUST_NEVER_CHANGE"
	// The value can change within stated constraints.
	CfnContextMutability_CHANGE_WITH_CONSTRAINTS CfnContextMutability = "CHANGE_WITH_CONSTRAINTS"
	// The value requires review before it changes.
	CfnContextMutability_REVIEW_REQUIRED CfnContextMutability = "REVIEW_REQUIRED"
	// The value can change without coordination.
	CfnContextMutability_FREE_TO_TUNE CfnContextMutability = "FREE_TO_TUNE"
)

