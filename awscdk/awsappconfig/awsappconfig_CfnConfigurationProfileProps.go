package awsappconfig


// Properties for defining a `CfnConfigurationProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationProfileProps := &cfnConfigurationProfileProps{
//   	applicationId: jsii.String("applicationId"),
//   	locationUri: jsii.String("locationUri"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	retrievalRoleArn: jsii.String("retrievalRoleArn"),
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   	validators: []interface{}{
//   		&validatorsProperty{
//   			content: jsii.String("content"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnConfigurationProfileProps struct {
	// The application ID.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// A URI to locate the configuration.
	//
	// You can specify the AWS AppConfig hosted configuration store, Systems Manager (SSM) document, an SSM Parameter Store parameter, or an Amazon S3 object. For the hosted configuration store and for feature flags, specify `hosted` . For an SSM document, specify either the document name in the format `ssm-document://<Document_name>` or the Amazon Resource Name (ARN). For a parameter, specify either the parameter name in the format `ssm-parameter://<Parameter_name>` or the ARN. For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>` . Here is an example: `s3://my-bucket/my-app/us-east-1/my-config.json`
	LocationUri *string `field:"required" json:"locationUri" yaml:"locationUri"`
	// A name for the configuration profile.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the configuration profile.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of an IAM role with permission to access the configuration at the specified `LocationUri` .
	//
	// > A retrieval role ARN is not required for configurations stored in the AWS AppConfig hosted configuration store. It is required for all other sources that store your configuration.
	RetrievalRoleArn *string `field:"optional" json:"retrievalRoleArn" yaml:"retrievalRoleArn"`
	// Metadata to assign to the configuration profile.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*CfnConfigurationProfile_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
	// The type of configurations contained in the profile.
	//
	// AWS AppConfig supports `feature flags` and `freeform` configurations. We recommend you create feature flag configurations to enable or disable new features and freeform configurations to distribute configurations to an application. When calling this API, enter one of the following values for `Type` :
	//
	// `AWS.AppConfig.FeatureFlags`
	//
	// `AWS.Freeform`
	Type *string `field:"optional" json:"type" yaml:"type"`
	// A list of methods for validating the configuration.
	Validators interface{} `field:"optional" json:"validators" yaml:"validators"`
}

