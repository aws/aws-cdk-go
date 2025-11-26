package previewawsdmsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMigrationProjectPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMigrationProjectMixinProps := &CfnMigrationProjectMixinProps{
//   	Description: jsii.String("description"),
//   	InstanceProfileArn: jsii.String("instanceProfileArn"),
//   	InstanceProfileIdentifier: jsii.String("instanceProfileIdentifier"),
//   	InstanceProfileName: jsii.String("instanceProfileName"),
//   	MigrationProjectCreationTime: jsii.String("migrationProjectCreationTime"),
//   	MigrationProjectIdentifier: jsii.String("migrationProjectIdentifier"),
//   	MigrationProjectName: jsii.String("migrationProjectName"),
//   	SchemaConversionApplicationAttributes: &SchemaConversionApplicationAttributesProperty{
//   		S3BucketPath: jsii.String("s3BucketPath"),
//   		S3BucketRoleArn: jsii.String("s3BucketRoleArn"),
//   	},
//   	SourceDataProviderDescriptors: []interface{}{
//   		&DataProviderDescriptorProperty{
//   			DataProviderArn: jsii.String("dataProviderArn"),
//   			DataProviderIdentifier: jsii.String("dataProviderIdentifier"),
//   			DataProviderName: jsii.String("dataProviderName"),
//   			SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   			SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetDataProviderDescriptors: []interface{}{
//   		&DataProviderDescriptorProperty{
//   			DataProviderArn: jsii.String("dataProviderArn"),
//   			DataProviderIdentifier: jsii.String("dataProviderIdentifier"),
//   			DataProviderName: jsii.String("dataProviderName"),
//   			SecretsManagerAccessRoleArn: jsii.String("secretsManagerAccessRoleArn"),
//   			SecretsManagerSecretId: jsii.String("secretsManagerSecretId"),
//   		},
//   	},
//   	TransformationRules: jsii.String("transformationRules"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html
//
type CfnMigrationProjectMixinProps struct {
	// A user-friendly description of the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the instance profile for your migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-instanceprofilearn
	//
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// The identifier of the instance profile for your migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-instanceprofileidentifier
	//
	InstanceProfileIdentifier *string `field:"optional" json:"instanceProfileIdentifier" yaml:"instanceProfileIdentifier"`
	// The name of the associated instance profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-instanceprofilename
	//
	InstanceProfileName *string `field:"optional" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The property describes a creating time of the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-migrationprojectcreationtime
	//
	// Deprecated: this property has been deprecated.
	MigrationProjectCreationTime *string `field:"optional" json:"migrationProjectCreationTime" yaml:"migrationProjectCreationTime"`
	// The identifier of the migration project.
	//
	// Identifiers must begin with a letter and must contain only ASCII letters, digits, and hyphens. They can't end with a hyphen, or contain two consecutive hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-migrationprojectidentifier
	//
	MigrationProjectIdentifier *string `field:"optional" json:"migrationProjectIdentifier" yaml:"migrationProjectIdentifier"`
	// The name of the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-migrationprojectname
	//
	MigrationProjectName *string `field:"optional" json:"migrationProjectName" yaml:"migrationProjectName"`
	// The schema conversion application attributes, including the Amazon S3 bucket name and Amazon S3 role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-schemaconversionapplicationattributes
	//
	SchemaConversionApplicationAttributes interface{} `field:"optional" json:"schemaConversionApplicationAttributes" yaml:"schemaConversionApplicationAttributes"`
	// Information about the source data provider, including the name or ARN, and AWS Secrets Manager parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-sourcedataproviderdescriptors
	//
	SourceDataProviderDescriptors interface{} `field:"optional" json:"sourceDataProviderDescriptors" yaml:"sourceDataProviderDescriptors"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Information about the target data provider, including the name or ARN, and AWS Secrets Manager parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-targetdataproviderdescriptors
	//
	TargetDataProviderDescriptors interface{} `field:"optional" json:"targetDataProviderDescriptors" yaml:"targetDataProviderDescriptors"`
	// The settings in JSON format for migration rules.
	//
	// Migration rules make it possible for you to change the object names according to the rules that you specify. For example, you can change an object name to lowercase or uppercase, add or remove a prefix or suffix, or rename objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-transformationrules
	//
	TransformationRules *string `field:"optional" json:"transformationRules" yaml:"transformationRules"`
}

