package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnConfigurationProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationProfileProps := &CfnConfigurationProfileProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	LocationUri: jsii.String("locationUri"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DeletionProtectionCheck: jsii.String("deletionProtectionCheck"),
//   	Description: jsii.String("description"),
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	RetrievalRoleArn: jsii.String("retrievalRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	Validators: []interface{}{
//   		&ValidatorsProperty{
//   			Content: jsii.String("content"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html
//
type CfnConfigurationProfileProps struct {
	// The application ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-applicationid
	//
	ApplicationId interface{} `field:"required" json:"applicationId" yaml:"applicationId"`
	// A URI to locate the configuration. You can specify the following:.
	//
	// - For the AWS AppConfig hosted configuration store and for feature flags, specify `hosted` .
	// - For an AWS Systems Manager Parameter Store parameter, specify either the parameter name in the format `ssm-parameter://<parameter name>` or the ARN.
	// - For an AWS CodePipeline pipeline, specify the URI in the following format: `codepipeline` ://<pipeline name>.
	// - For an AWS Secrets Manager secret, specify the URI in the following format: `secretsmanager` ://<secret name>.
	// - For an Amazon S3 object, specify the URI in the following format: `s3://<bucket>/<objectKey>` . Here is an example: `s3://amzn-s3-demo-bucket/my-app/us-east-1/my-config.json`
	// - For an SSM document, specify either the document name in the format `ssm-document://<document name>` or the Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-locationuri
	//
	LocationUri *string `field:"required" json:"locationUri" yaml:"locationUri"`
	// A name for the configuration profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A parameter to configure deletion protection.
	//
	// Deletion protection prevents a user from deleting a configuration profile if your application has called either [GetLatestConfiguration](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_appconfigdata_GetLatestConfiguration.html) or [GetConfiguration](https://docs.aws.amazon.com/appconfig/2019-10-09/APIReference/API_GetConfiguration.html) for the configuration profile during the specified interval.
	//
	// This parameter supports the following values:
	//
	// - `BYPASS` : Instructs AWS AppConfig to bypass the deletion protection check and delete a configuration profile even if deletion protection would have otherwise prevented it.
	// - `APPLY` : Instructs the deletion protection check to run, even if deletion protection is disabled at the account level. `APPLY` also forces the deletion protection check to run against resources created in the past hour, which are normally excluded from deletion protection checks.
	// - `ACCOUNT_DEFAULT` : The default setting, which instructs AWS AppConfig to implement the deletion protection value specified in the `UpdateAccountSettings` API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-deletionprotectioncheck
	//
	DeletionProtectionCheck *string `field:"optional" json:"deletionProtectionCheck" yaml:"deletionProtectionCheck"`
	// A description of the configuration profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-kmskeyidentifier
	//
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The ARN of an IAM role with permission to access the configuration at the specified `LocationUri` .
	//
	// > A retrieval role ARN is not required for configurations stored in AWS CodePipeline or the AWS AppConfig hosted configuration store. It is required for all other sources that store your configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-retrievalrolearn
	//
	RetrievalRoleArn interface{} `field:"optional" json:"retrievalRoleArn" yaml:"retrievalRoleArn"`
	// Metadata to assign to the configuration profile.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of configurations contained in the profile.
	//
	// AWS AppConfig supports `feature flags` and `freeform` configurations. We recommend you create feature flag configurations to enable or disable new features and freeform configurations to distribute configurations to an application. When calling this API, enter one of the following values for `Type` :
	//
	// `AWS.AppConfig.FeatureFlags`
	//
	// `AWS.Freeform`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// A list of methods for validating the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-configurationprofile.html#cfn-appconfig-configurationprofile-validators
	//
	Validators interface{} `field:"optional" json:"validators" yaml:"validators"`
}

