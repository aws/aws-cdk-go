package awsappconfig


// A value such as an Amazon Resource Name (ARN) or an Amazon Simple Notification Service topic entered in an extension when invoked.
//
// Parameter values are specified in an extension association. For more information about extensions, see [Working with AWS AppConfig extensions](https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html) in the *AWS AppConfig User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterProperty := &ParameterProperty{
//   	Required: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-parameter.html
//
type CfnExtension_ParameterProperty struct {
	// A parameter value must be specified in the extension association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-parameter.html#cfn-appconfig-extension-parameter-required
	//
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// Information about the parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-parameter.html#cfn-appconfig-extension-parameter-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

