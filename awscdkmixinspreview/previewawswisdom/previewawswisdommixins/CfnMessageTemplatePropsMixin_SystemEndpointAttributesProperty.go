package previewawswisdommixins


// The system endpoint attributes that are used with the message template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   systemEndpointAttributesProperty := &SystemEndpointAttributesProperty{
//   	Address: jsii.String("address"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-systemendpointattributes.html
//
type CfnMessageTemplatePropsMixin_SystemEndpointAttributesProperty struct {
	// The customer's phone number if used with `customerEndpoint` , or the number the customer dialed to call your contact center if used with `systemEndpoint` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-messagetemplate-systemendpointattributes.html#cfn-wisdom-messagetemplate-systemendpointattributes-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
}

