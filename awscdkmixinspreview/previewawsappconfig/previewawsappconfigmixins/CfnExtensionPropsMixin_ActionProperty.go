package previewawsappconfigmixins


// The actions defined in the extension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   actionProperty := &ActionProperty{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Uri: jsii.String("uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html
//
type CfnExtensionPropsMixin_ActionProperty struct {
	// Information about actions defined in the extension.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The extension name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An Amazon Resource Name (ARN) for an AWS Identity and Access Management assume role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The extension URI associated to the action point in the extension definition.
	//
	// The URI can be an Amazon Resource Name (ARN) for one of the following: an AWS Lambda function, an Amazon Simple Queue Service queue, an Amazon Simple Notification Service topic, or the Amazon EventBridge default event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-uri
	//
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

