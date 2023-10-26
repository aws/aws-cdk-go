package awsappconfig


// An action defines the tasks that the extension performs during the AWS AppConfig workflow.
//
// Each action includes an action point such as `ON_CREATE_HOSTED_CONFIGURATION` , `PRE_DEPLOYMENT` , or `ON_DEPLOYMENT` . Each action also includes a name, a URI to an AWS Lambda function, and an Amazon Resource Name (ARN) for an AWS Identity and Access Management assume role. You specify the name, URI, and ARN for each *action point* defined in the extension. You can specify the following actions for an extension:
//
// - `PRE_CREATE_HOSTED_CONFIGURATION_VERSION`
// - `PRE_START_DEPLOYMENT`
// - `ON_DEPLOYMENT_START`
// - `ON_DEPLOYMENT_STEP`
// - `ON_DEPLOYMENT_BAKING`
// - `ON_DEPLOYMENT_COMPLETE`
// - `ON_DEPLOYMENT_ROLLED_BACK`.
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
	// The action name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The extension URI associated to the action point in the extension definition.
	//
	// The URI can be an Amazon Resource Name (ARN) for one of the following: an AWS Lambda function, an Amazon Simple Queue Service queue, an Amazon Simple Notification Service topic, or the Amazon EventBridge default event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-uri
	//
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Information about the action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An Amazon Resource Name (ARN) for an AWS Identity and Access Management assume role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-extension-action.html#cfn-appconfig-extension-action-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

