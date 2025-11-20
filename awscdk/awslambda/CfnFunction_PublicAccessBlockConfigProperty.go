package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicAccessBlockConfigProperty := &PublicAccessBlockConfigProperty{
//   	BlockPublicPolicy: jsii.Boolean(false),
//   	RestrictPublicResource: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-publicaccessblockconfig.html
//
type CfnFunction_PublicAccessBlockConfigProperty struct {
	// Block Public Policy from being attached to a function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-publicaccessblockconfig.html#cfn-lambda-function-publicaccessblockconfig-blockpublicpolicy
	//
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Restrict public resource access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-publicaccessblockconfig.html#cfn-lambda-function-publicaccessblockconfig-restrictpublicresource
	//
	RestrictPublicResource interface{} `field:"optional" json:"restrictPublicResource" yaml:"restrictPublicResource"`
}

