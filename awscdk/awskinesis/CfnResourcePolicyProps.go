package awskinesis


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-resourcepolicy.html
//
type CfnResourcePolicyProps struct {
	// Returns the Amazon Resource Name (ARN) of the resource-based policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-resourcepolicy.html#cfn-kinesis-resourcepolicy-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// This is the description for the resource policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-resourcepolicy.html#cfn-kinesis-resourcepolicy-resourcepolicy
	//
	ResourcePolicy interface{} `field:"required" json:"resourcePolicy" yaml:"resourcePolicy"`
}

