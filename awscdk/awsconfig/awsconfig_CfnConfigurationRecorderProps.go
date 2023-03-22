package awsconfig


// Properties for defining a `CfnConfigurationRecorder`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationRecorderProps := &cfnConfigurationRecorderProps{
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	recordingGroup: &recordingGroupProperty{
//   		allSupported: jsii.Boolean(false),
//   		includeGlobalResourceTypes: jsii.Boolean(false),
//   		resourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   	},
//   }
//
type CfnConfigurationRecorderProps struct {
	// The Amazon Resource Name (ARN) of the IAM (IAM) role that is used to make read or write requests to the delivery channel that you specify and to get configuration details for supported AWS resources.
	//
	// For more information, see [Permissions for the IAM Role Assigned](https://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) to AWS Config in the AWS Config Developer Guide.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A name for the configuration recorder.
	//
	// If you don't specify a name, AWS CloudFormation CloudFormation generates a unique physical ID and uses that ID for the configuration recorder name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > After you create a configuration recorder, you cannot rename it. If you don't want a name that AWS CloudFormation generates, specify a value for this property.
	//
	// Updates are not supported.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicates whether to record configurations for all supported resources or for a list of resource types.
	//
	// The resource types that you list must be supported by AWS Config .
	RecordingGroup interface{} `field:"optional" json:"recordingGroup" yaml:"recordingGroup"`
}

