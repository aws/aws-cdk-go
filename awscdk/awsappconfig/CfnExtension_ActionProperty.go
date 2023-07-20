package awsappconfig


// An action for an extension to take at a specific action point.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   actionProperty := &ActionProperty{
//   	Name: jsii.String("name"),
//   	Uri: jsii.String("uri"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html
//
type CfnExtension_ActionProperty struct {
	// The name of the extension action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URI of the extension action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-uri
	//
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The description of the extension Action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the role for invoking the extension action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

