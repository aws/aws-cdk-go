package mixinsawslogs


// Properties for CfnAccountPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccountPolicyMixinProps := &CfnAccountPolicyMixinProps{
//   	PolicyDocument: jsii.String("policyDocument"),
//   	PolicyName: jsii.String("policyName"),
//   	PolicyType: jsii.String("policyType"),
//   	Scope: jsii.String("scope"),
//   	SelectionCriteria: jsii.String("selectionCriteria"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html
//
type CfnAccountPolicyMixinProps struct {
	// Specify the policy, in JSON.
	//
	// *Data protection policy*
	//
	// A data protection policy must include two JSON blocks:
	//
	// - The first block must include both a `DataIdentifer` array and an `Operation` property with an `Audit` action. The `DataIdentifer` array lists the types of sensitive data that you want to mask. For more information about the available options, see [Types of data that you can mask](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data-types.html) .
	//
	// The `Operation` property with an `Audit` action is required to find the sensitive data terms. This `Audit` action must contain a `FindingsDestination` object. You can optionally use that `FindingsDestination` object to list one or more destinations to send audit findings to. If you specify destinations such as log groups, Firehose streams, and S3 buckets, they must already exist.
	// - The second block must include both a `DataIdentifer` array and an `Operation` property with an `Deidentify` action. The `DataIdentifer` array must exactly match the `DataIdentifer` array in the first block of the policy.
	//
	// The `Operation` property with the `Deidentify` action is what actually masks the data, and it must contain the `"MaskConfig": {}` object. The `"MaskConfig": {}` object must be empty.
	//
	// > The contents of the two `DataIdentifer` arrays must match exactly.
	//
	// In addition to the two JSON blocks, the `policyDocument` can also include `Name` , `Description` , and `Version` fields. The `Name` is different than the operation's `policyName` parameter, and is used as a dimension when CloudWatch Logs reports audit findings metrics to CloudWatch .
	//
	// The JSON specified in `policyDocument` can be up to 30,720 characters long.
	//
	// *Subscription filter policy*
	//
	// A subscription filter policy can include the following attributes in a JSON block:
	//
	// - *DestinationArn* The ARN of the destination to deliver log events to. Supported destinations are:
	//
	// - An Kinesis Data Streams data stream in the same account as the subscription policy, for same-account delivery.
	// - An Firehose data stream in the same account as the subscription policy, for same-account delivery.
	// - A Lambda function in the same account as the subscription policy, for same-account delivery.
	// - A logical destination in a different account created with [PutDestination](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_PutDestination.html) , for cross-account delivery. Kinesis Data Streams and Firehose are supported as logical destinations.
	// - *RoleArn* The ARN of an IAM role that grants CloudWatch Logs permissions to deliver ingested log events to the destination stream. You don't need to provide the ARN when you are working with a logical destination for cross-account delivery.
	// - *FilterPattern* A filter pattern for subscribing to a filtered stream of log events.
	// - *Distribution* The method used to distribute log data to the destination. By default, log data is grouped by log stream, but the grouping can be set to `Random` for a more even distribution. This property is only applicable when the destination is an Kinesis Data Streams data stream.
	//
	// *Field index policy*
	//
	// A field index filter policy can include the following attribute in a JSON block:
	//
	// - *Fields* The array of field indexes to create.
	//
	// The following is an example of an index policy document that creates two indexes, `RequestId` and `TransactionId` .
	//
	// `"policyDocument": "{ \"Fields\": [ \"RequestId\", \"TransactionId\" ] }"`
	//
	// *Transformer policy*
	//
	// A transformer policy must include one JSON block with the array of processors and their configurations. For more information about available processors, see [Processors that you can use](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch-Logs-Transformation.html#CloudWatch-Logs-Transformation-Processors) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-policydocument
	//
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// A name for the policy.
	//
	// This must be unique within the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-policyname
	//
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// The type of policy that you're creating or updating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-policytype
	//
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// Currently the only valid value for this parameter is `ALL` , which specifies that the policy applies to all log groups in the account.
	//
	// If you omit this parameter, the default of `ALL` is used. To scope down a subscription filter policy to a subset of log groups, use the `SelectionCriteria` parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Use this parameter to apply the new policy to a subset of log groups in the account.
	//
	// You need to specify `SelectionCriteria` only when you specify `SUBSCRIPTION_FILTER_POLICY` , `FIELD_INDEX_POLICY` or `TRANSFORMER_POLICY` for `PolicyType` .
	//
	// If `PolicyType` is `SUBSCRIPTION_FILTER_POLICY` , the only supported `SelectionCriteria` filter is `LogGroupName NOT IN []`
	//
	// If `PolicyType` is `FIELD_INDEX_POLICY` or `TRANSFORMER_POLICY` , the only supported `SelectionCriteria` filter is `LogGroupNamePrefix`
	//
	// The `SelectionCriteria` string can be up to 25KB in length. The length is determined by using its UTF-8 bytes.
	//
	// Using the `SelectionCriteria` parameter with `SUBSCRIPTION_FILTER_POLICY` is useful to help prevent infinite loops. For more information, see [Log recursion prevention](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Subscriptions-recursion-prevention.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-accountpolicy.html#cfn-logs-accountpolicy-selectioncriteria
	//
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
}

