package awscdk


// Template-level CloudFormation Metadata Context.
//
// Example:
//   var stack Stack
//
//
//   awscdk.CfnTemplateMetadataContext_Of(stack).Set(&CfnTemplateMetadataContextProps{
//   	Arch: jsii.String("SQS buffer -> Lambda -> DynamoDB"),
//   	Owner: jsii.String("order-processing-team"),
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-metadata.html#aws-attribute-metadata-context-schema
//
type CfnTemplateMetadataContextProps struct {
	// High-level architecture of the template.
	// Default: - no architecture overview.
	//
	Arch *string `field:"optional" json:"arch" yaml:"arch"`
	// Constraints that apply across the template.
	// Default: - no constraints.
	//
	Must *[]*string `field:"optional" json:"must" yaml:"must"`
	// Owner of the template.
	// Default: - no owner.
	//
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// References to supporting context.
	// Default: - no references.
	//
	Ref *[]*CfnContextRef `field:"optional" json:"ref" yaml:"ref"`
}

