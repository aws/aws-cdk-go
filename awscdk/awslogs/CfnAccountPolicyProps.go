package awslogs


// Properties for defining a `CfnAccountPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountPolicyProps := &CfnAccountPolicyProps{
//   	PolicyDocument: jsii.String("policyDocument"),
//   	PolicyName: jsii.String("policyName"),
//   	PolicyType: jsii.String("policyType"),
//
//   	// the properties below are optional
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html
//
type CfnAccountPolicyProps struct {
	// Specify the data protection policy, in JSON.
	//
	// This policy must include two JSON blocks:
	//
	// - The first block must include both a `DataIdentifer` array and an `Operation` property with an `Audit` action. The `DataIdentifer` array lists the types of sensitive data that you want to mask. For more information about the available options, see [Types of data that you can mask](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-types.html) .
	//
	// The `Operation` property with an `Audit` action is required to find the sensitive data terms. This `Audit` action must contain a `FindingsDestination` object. You can optionally use that `FindingsDestination` object to list one or more destinations to send audit findings to. If you specify destinations such as log groups, Kinesis Data Firehose streams, and S3 buckets, they must already exist.
	// - The second block must include both a `DataIdentifer` array and an `Operation` property with an `Deidentify` action. The `DataIdentifer` array must exactly match the `DataIdentifer` array in the first block of the policy.
	//
	// The `Operation` property with the `Deidentify` action is what actually masks the data, and it must contain the `"MaskConfig": {}` object. The `"MaskConfig": {}` object must be empty.
	//
	// > The contents of the two `DataIdentifer` arrays must match exactly.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-policydocument
	//
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// A name for the policy.
	//
	// This must be unique within the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-policyname
	//
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Currently the only valid value for this parameter is `DATA_PROTECTION_POLICY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-policytype
	//
	PolicyType *string `field:"required" json:"policyType" yaml:"policyType"`
	// Currently the only valid value for this parameter is `ALL` , which specifies that the data protection policy applies to all log groups in the account.
	//
	// If you omit this parameter, the default of `ALL` is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}

