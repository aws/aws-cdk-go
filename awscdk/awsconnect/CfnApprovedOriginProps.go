package awsconnect


// Properties for defining a `CfnApprovedOrigin`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApprovedOriginProps := &CfnApprovedOriginProps{
//   	InstanceId: jsii.String("instanceId"),
//   	Origin: jsii.String("origin"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-approvedorigin.html
//
type CfnApprovedOriginProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	//
	// *Minimum* : `1`
	//
	// *Maximum* : `100`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-approvedorigin.html#cfn-connect-approvedorigin-instanceid
	//
	InstanceId interface{} `field:"required" json:"instanceId" yaml:"instanceId"`
	// Domain name to be added to the allow-list of the instance.
	//
	// *Maximum* : `267`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-approvedorigin.html#cfn-connect-approvedorigin-origin
	//
	Origin *string `field:"required" json:"origin" yaml:"origin"`
}

