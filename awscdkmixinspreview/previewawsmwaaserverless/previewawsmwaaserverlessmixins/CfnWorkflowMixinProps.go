package previewawsmwaaserverlessmixins


// Properties for CfnWorkflowPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkflowMixinProps := &CfnWorkflowMixinProps{
//   	DefinitionS3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		ObjectKey: jsii.String("objectKey"),
//   		VersionId: jsii.String("versionId"),
//   	},
//   	Description: jsii.String("description"),
//   	EncryptionConfiguration: &EncryptionConfigurationProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		Type: jsii.String("type"),
//   	},
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   	},
//   	Name: jsii.String("name"),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TriggerMode: jsii.String("triggerMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html
//
type CfnWorkflowMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html#cfn-mwaaserverless-workflow-definitions3location
	//
	DefinitionS3Location interface{} `field:"optional" json:"definitionS3Location" yaml:"definitionS3Location"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html#cfn-mwaaserverless-workflow-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html#cfn-mwaaserverless-workflow-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html#cfn-mwaaserverless-workflow-loggingconfiguration
	//
	LoggingConfiguration interface{} `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html#cfn-mwaaserverless-workflow-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html#cfn-mwaaserverless-workflow-networkconfiguration
	//
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html#cfn-mwaaserverless-workflow-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A map of key-value pairs to be applied as tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html#cfn-mwaaserverless-workflow-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaaserverless-workflow.html#cfn-mwaaserverless-workflow-triggermode
	//
	TriggerMode *string `field:"optional" json:"triggerMode" yaml:"triggerMode"`
}

