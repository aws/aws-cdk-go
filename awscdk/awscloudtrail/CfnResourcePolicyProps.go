package awscloudtrail


// Properties for defining a `CfnResourcePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var resourcePolicy interface{}
//
//   cfnResourcePolicyProps := &CfnResourcePolicyProps{
//   	ResourceArn: jsii.String("resourceArn"),
//   	ResourcePolicy: resourcePolicy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-resourcepolicy.html
//
type CfnResourcePolicyProps struct {
	// The Amazon Resource Name (ARN) of the CloudTrail channel attached to the resource-based policy.
	//
	// The following is the format of a resource ARN: `arn:aws:cloudtrail:us-east-2:123456789012:channel/MyChannel` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-resourcepolicy.html#cfn-cloudtrail-resourcepolicy-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// A JSON-formatted string for an AWS resource-based policy.
	//
	// The following are requirements for the resource policy:
	//
	// - Contains only one action: cloudtrail-data:PutAuditEvents
	// - Contains at least one statement. The policy can have a maximum of 20 statements.
	// - Each statement contains at least one principal. A statement can have a maximum of 50 principals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-resourcepolicy.html#cfn-cloudtrail-resourcepolicy-resourcepolicy
	//
	ResourcePolicy interface{} `field:"required" json:"resourcePolicy" yaml:"resourcePolicy"`
}

