package awscdk


// Confidence in the accuracy of context.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-metadata.html#aws-attribute-metadata-context-schema
//
type CfnContextTrustConfidence string

const (
	// High confidence.
	CfnContextTrustConfidence_HIGH CfnContextTrustConfidence = "HIGH"
	// Medium confidence.
	CfnContextTrustConfidence_MEDIUM CfnContextTrustConfidence = "MEDIUM"
	// Low confidence.
	CfnContextTrustConfidence_LOW CfnContextTrustConfidence = "LOW"
)

