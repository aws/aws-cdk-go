package awsssm


// Additional properties for looking up an existing StringParameter.
//
// Example:
//   stringValue := ssm.StringParameter_ValueFromLookup(this, jsii.String("/My/Public/Parameter"), undefined, &StringParameterLookupOptions{
//   	AdditionalCacheKey: this.Node.path,
//   })
//
type StringParameterLookupOptions struct {
	// Adds an additional discriminator to the `cdk.context.json` cache key.
	// Default: - no additional cache key.
	//
	AdditionalCacheKey *string `field:"optional" json:"additionalCacheKey" yaml:"additionalCacheKey"`
}

