package mixinsawsqbusiness


// Configuration information about chat response personalization.
//
// For more information, see [Personalizing chat responses](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/personalizing-chat-responses.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   personalizationConfigurationProperty := &PersonalizationConfigurationProperty{
//   	PersonalizationControlMode: jsii.String("personalizationControlMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-personalizationconfiguration.html
//
type CfnApplicationPropsMixin_PersonalizationConfigurationProperty struct {
	// An option to allow Amazon Q Business to customize chat responses using user specific metadata—specifically, location and job information—in your IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-application-personalizationconfiguration.html#cfn-qbusiness-application-personalizationconfiguration-personalizationcontrolmode
	//
	PersonalizationControlMode *string `field:"optional" json:"personalizationControlMode" yaml:"personalizationControlMode"`
}

