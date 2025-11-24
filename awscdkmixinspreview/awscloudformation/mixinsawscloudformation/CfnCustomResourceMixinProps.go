package mixinsawscloudformation


// Properties for CfnCustomResourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomResourceMixinProps := &CfnCustomResourceMixinProps{
//   	ServiceTimeout: jsii.Number(123),
//   	ServiceToken: jsii.String("serviceToken"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-customresource.html
//
type CfnCustomResourceMixinProps struct {
	// The maximum time, in seconds, that can elapse before a custom resource operation times out.
	//
	// The value must be an integer from 1 to 3600. The default value is 3600 seconds (1 hour).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-customresource.html#cfn-cloudformation-customresource-servicetimeout
	//
	ServiceTimeout *float64 `field:"optional" json:"serviceTimeout" yaml:"serviceTimeout"`
	// The service token, such as an Amazon  topic ARN or Lambda function ARN.
	//
	// The service token must be from the same Region as the stack.
	//
	// Updates aren't supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-customresource.html#cfn-cloudformation-customresource-servicetoken
	//
	ServiceToken *string `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}

