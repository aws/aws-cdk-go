package awslightsail


// Properties for CfnContactMethodPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnContactMethodMixinProps := &CfnContactMethodMixinProps{
//   	ContactEndpoint: jsii.String("contactEndpoint"),
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-contactmethod.html
//
type CfnContactMethodMixinProps struct {
	// The destination of the contact method, such as an email address or a mobile phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-contactmethod.html#cfn-lightsail-contactmethod-contactendpoint
	//
	ContactEndpoint *string `field:"optional" json:"contactEndpoint" yaml:"contactEndpoint"`
	// The protocol of the contact method, such as Email or SMS (text messaging).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-contactmethod.html#cfn-lightsail-contactmethod-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

