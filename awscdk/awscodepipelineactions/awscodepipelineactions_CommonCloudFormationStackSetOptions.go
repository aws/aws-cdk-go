package awscodepipelineactions


// Options in common between both StackSet actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   commonCloudFormationStackSetOptions := &commonCloudFormationStackSetOptions{
//   	failureTolerancePercentage: jsii.Number(123),
//   	maxAccountConcurrencyPercentage: jsii.Number(123),
//   	stackSetRegion: jsii.String("stackSetRegion"),
//   }
//
type CommonCloudFormationStackSetOptions struct {
	// The percentage of accounts per Region for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	//
	// If
	// the operation is stopped in a Region, AWS CloudFormation doesn't attempt the operation in subsequent Regions. When calculating the number
	// of accounts based on the specified percentage, AWS CloudFormation rounds down to the next whole number.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified
	// percentage, AWS CloudFormation rounds down to the next whole number. If rounding down would result in zero, AWS CloudFormation sets the number as
	// one instead. Although you use this setting to specify the maximum, for large deployments the actual number of accounts acted upon concurrently
	// may be lower due to service throttling.
	MaxAccountConcurrencyPercentage *float64 `field:"optional" json:"maxAccountConcurrencyPercentage" yaml:"maxAccountConcurrencyPercentage"`
	// The AWS Region the StackSet is in.
	//
	// Note that a cross-region Pipeline requires replication buckets to function correctly.
	// You can provide their names with the `PipelineProps.crossRegionReplicationBuckets` property.
	// If you don't, the CodePipeline Construct will create new Stacks in your CDK app containing those buckets,
	// that you will need to `cdk deploy` before deploying the main, Pipeline-containing Stack.
	StackSetRegion *string `field:"optional" json:"stackSetRegion" yaml:"stackSetRegion"`
}

