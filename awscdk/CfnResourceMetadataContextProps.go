package awscdk


// Resource-level CloudFormation Metadata Context.
//
// Example:
//   var queue Queue
//
//
//   awscdk.CfnResourceMetadataContext_Of(queue).Set(&CfnResourceMetadataContextProps{
//   	Why: jsii.String("buffers order events for asynchronous processing"),
//   	Must: []*string{
//   		jsii.String("VisibilityTimeout must be at least six times the consumer timeout"),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-attribute-metadata.html#aws-attribute-metadata-context-schema
//
type CfnResourceMetadataContextProps struct {
	// Dependencies of the resource.
	// Default: - no dependencies.
	//
	Deps *[]*string `field:"optional" json:"deps" yaml:"deps"`
	// Constraints that must remain true.
	// Default: - no constraints.
	//
	Must *[]*string `field:"optional" json:"must" yaml:"must"`
	// Mutability by CloudFormation property name.
	// Default: - no property mutability.
	//
	Mutability *map[string]CfnContextMutability `field:"optional" json:"mutability" yaml:"mutability"`
	// Default mutability of the resource.
	// Default: - no default mutability.
	//
	Mutable CfnContextMutability `field:"optional" json:"mutable" yaml:"mutable"`
	// Provenance and confidence for the context.
	// Default: - no trust information.
	//
	Trust *CfnContextTrust `field:"optional" json:"trust" yaml:"trust"`
	// Rationale for the resource.
	// Default: - no rationale.
	//
	Why *string `field:"optional" json:"why" yaml:"why"`
}

