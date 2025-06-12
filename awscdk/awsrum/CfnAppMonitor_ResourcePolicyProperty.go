package awsrum


// Use this structure to assign a resource-based policy to a CloudWatch RUM app monitor to control access to it.
//
// Each app monitor can have one resource-based policy. The maximum size of the policy is 4 KB. To learn more about using resource policies with RUM, see [Using resource-based policies with CloudWatch RUM](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-resource-policies.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePolicyProperty := &ResourcePolicyProperty{
//   	PolicyDocument: jsii.String("policyDocument"),
//
//   	// the properties below are optional
//   	PolicyRevisionId: jsii.String("policyRevisionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-resourcepolicy.html
//
type CfnAppMonitor_ResourcePolicyProperty struct {
	// The JSON to use as the resource policy.
	//
	// The document can be up to 4 KB in size. For more information about the contents and syntax for this policy, see [Using resource-based policies with CloudWatch RUM](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-resource-policies.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-resourcepolicy.html#cfn-rum-appmonitor-resourcepolicy-policydocument
	//
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// A string value that you can use to conditionally update your policy.
	//
	// You can provide the revision ID of your existing policy to make mutating requests against that policy.
	//
	// When you assign a policy revision ID, then later requests about that policy will be rejected with an `InvalidPolicyRevisionIdException` error if they don't provide the correct current revision ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rum-appmonitor-resourcepolicy.html#cfn-rum-appmonitor-resourcepolicy-policyrevisionid
	//
	PolicyRevisionId *string `field:"optional" json:"policyRevisionId" yaml:"policyRevisionId"`
}

