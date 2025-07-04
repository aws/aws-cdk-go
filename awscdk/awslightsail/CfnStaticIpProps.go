package awslightsail


// Properties for defining a `CfnStaticIp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStaticIpProps := &CfnStaticIpProps{
//   	StaticIpName: jsii.String("staticIpName"),
//
//   	// the properties below are optional
//   	AttachedTo: jsii.String("attachedTo"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-staticip.html
//
type CfnStaticIpProps struct {
	// The name of the static IP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-staticip.html#cfn-lightsail-staticip-staticipname
	//
	StaticIpName *string `field:"required" json:"staticIpName" yaml:"staticIpName"`
	// The instance that the static IP is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-staticip.html#cfn-lightsail-staticip-attachedto
	//
	AttachedTo *string `field:"optional" json:"attachedTo" yaml:"attachedTo"`
}

