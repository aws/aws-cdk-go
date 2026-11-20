package awscdk


// How context was produced.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-metadata.html#aws-attribute-metadata-context-schema
//
type CfnContextTrustSource string

const (
	// Human-authored or human-confirmed context.
	CfnContextTrustSource_AUTHORED CfnContextTrustSource = "AUTHORED"
	// Context derived from a code comment.
	CfnContextTrustSource_COMMENT CfnContextTrustSource = "COMMENT"
	// Context derived from a commit message.
	CfnContextTrustSource_COMMIT CfnContextTrustSource = "COMMIT"
	// Context inferred by a tool.
	CfnContextTrustSource_INFER CfnContextTrustSource = "INFER"
)

