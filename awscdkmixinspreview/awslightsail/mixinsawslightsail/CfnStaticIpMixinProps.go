package mixinsawslightsail


// Properties for CfnStaticIpPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStaticIpMixinProps := &CfnStaticIpMixinProps{
//   	AttachedTo: jsii.String("attachedTo"),
//   	StaticIpName: jsii.String("staticIpName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-staticip.html
//
type CfnStaticIpMixinProps struct {
	// The instance that the static IP is attached to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-staticip.html#cfn-lightsail-staticip-attachedto
	//
	AttachedTo *string `field:"optional" json:"attachedTo" yaml:"attachedTo"`
	// The name of the static IP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-staticip.html#cfn-lightsail-staticip-staticipname
	//
	StaticIpName *string `field:"optional" json:"staticIpName" yaml:"staticIpName"`
}

