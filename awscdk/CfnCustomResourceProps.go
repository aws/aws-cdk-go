package awscdk


// Properties for defining a `CfnCustomResource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomResourceProps := &CfnCustomResourceProps{
//   	ServiceToken: jsii.String("serviceToken"),
//
//   	// the properties below are optional
//   	ServiceTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-customresource.html
//
type CfnCustomResourceProps struct {
	// The service token, such as an Amazon SNS topic ARN or Lambda function ARN.
	//
	// The service token must be from the same Region as the stack.
	//
	// Updates aren't supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-customresource.html#cfn-cloudformation-customresource-servicetoken
	//
	ServiceToken *string `field:"required" json:"serviceToken" yaml:"serviceToken"`
	// The maximum time, in seconds, that can elapse before a custom resource operation times out.
	//
	// The value must be an integer from 1 to 3600. The default value is 3600 seconds (1 hour).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-customresource.html#cfn-cloudformation-customresource-servicetimeout
	//
	ServiceTimeout *float64 `field:"optional" json:"serviceTimeout" yaml:"serviceTimeout"`
}

