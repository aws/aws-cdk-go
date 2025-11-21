package awscloudformation


// The user-specified preferences for how CloudFormation performs a StackSet operation.
//
// For more information on maximum concurrent accounts and failure tolerance, see [StackSet operation options](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html#stackset-ops-options) in the *CloudFormation User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   operationPreferencesProperty := &OperationPreferencesProperty{
//   	ConcurrencyMode: jsii.String("concurrencyMode"),
//   	FailureToleranceCount: jsii.Number(123),
//   	FailureTolerancePercentage: jsii.Number(123),
//   	MaxConcurrentCount: jsii.Number(123),
//   	MaxConcurrentPercentage: jsii.Number(123),
//   	RegionConcurrencyType: jsii.String("regionConcurrencyType"),
//   	RegionOrder: []*string{
//   		jsii.String("regionOrder"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html
//
type CfnStackSet_OperationPreferencesProperty struct {
	// Specifies how the concurrency level behaves during the operation execution.
	//
	// - `STRICT_FAILURE_TOLERANCE` : This option dynamically lowers the concurrency level to ensure the number of failed accounts never exceeds the value of `FailureToleranceCount` +1. The initial actual concurrency is set to the lower of either the value of the `MaxConcurrentCount` , or the value of `FailureToleranceCount` +1. The actual concurrency is then reduced proportionally by the number of failures. This is the default behavior.
	//
	// If failure tolerance or Maximum concurrent accounts are set to percentages, the behavior is similar.
	// - `SOFT_FAILURE_TOLERANCE` : This option decouples `FailureToleranceCount` from the actual concurrency. This allows StackSet operations to run at the concurrency level set by the `MaxConcurrentCount` value, or `MaxConcurrentPercentage` , regardless of the number of failures.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-concurrencymode
	//
	ConcurrencyMode *string `field:"optional" json:"concurrencyMode" yaml:"concurrencyMode"`
	// The number of accounts per Region this operation can fail in before CloudFormation stops the operation in that Region.
	//
	// If the operation is stopped in a Region, CloudFormation doesn't attempt the operation in any subsequent Regions.
	//
	// Conditional: You must specify either `FailureToleranceCount` or `FailureTolerancePercentage` (but not both).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-failuretolerancecount
	//
	FailureToleranceCount *float64 `field:"optional" json:"failureToleranceCount" yaml:"failureToleranceCount"`
	// The percentage of accounts per Region this stack operation can fail in before CloudFormation stops the operation in that Region.
	//
	// If the operation is stopped in a Region, CloudFormation doesn't attempt the operation in any subsequent Regions.
	//
	// When calculating the number of accounts based on the specified percentage, CloudFormation rounds *down* to the next whole number.
	//
	// Conditional: You must specify either `FailureToleranceCount` or `FailureTolerancePercentage` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-failuretolerancepercentage
	//
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// The maximum number of accounts in which to perform this operation at one time.
	//
	// This is dependent on the value of `FailureToleranceCount` . `MaxConcurrentCount` is at most one more than the `FailureToleranceCount` .
	//
	// Note that this setting lets you specify the *maximum* for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
	//
	// Conditional: You must specify either `MaxConcurrentCount` or `MaxConcurrentPercentage` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-maxconcurrentcount
	//
	MaxConcurrentCount *float64 `field:"optional" json:"maxConcurrentCount" yaml:"maxConcurrentCount"`
	// The maximum percentage of accounts in which to perform this operation at one time.
	//
	// When calculating the number of accounts based on the specified percentage, CloudFormation rounds down to the next whole number. This is true except in cases where rounding down would result is zero. In this case, CloudFormation sets the number as one instead.
	//
	// Note that this setting lets you specify the *maximum* for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.
	//
	// Conditional: You must specify either `MaxConcurrentCount` or `MaxConcurrentPercentage` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-maxconcurrentpercentage
	//
	MaxConcurrentPercentage *float64 `field:"optional" json:"maxConcurrentPercentage" yaml:"maxConcurrentPercentage"`
	// The concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-regionconcurrencytype
	//
	RegionConcurrencyType *string `field:"optional" json:"regionConcurrencyType" yaml:"regionConcurrencyType"`
	// The order of the Regions where you want to perform the stack operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-operationpreferences.html#cfn-cloudformation-stackset-operationpreferences-regionorder
	//
	RegionOrder *[]*string `field:"optional" json:"regionOrder" yaml:"regionOrder"`
}

