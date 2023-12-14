package awsdms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMigrationProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMigrationProjectProps := &CfnMigrationProjectProps{
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
//   	Tags: []cfnTag{
//   		&cfnTag{
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
type CfnMigrationProjectProps struct {
	// The optional description of the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The property describes an instance profile arn for the migration project.
	//
	// For read.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-instanceprofilearn
	//
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// The property describes an instance profile identifier for the migration project.
	//
	// For create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-instanceprofileidentifier
	//
	InstanceProfileIdentifier *string `field:"optional" json:"instanceProfileIdentifier" yaml:"instanceProfileIdentifier"`
	// The property describes an instance profile name for the migration project.
	//
	// For read.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-instanceprofilename
	//
	InstanceProfileName *string `field:"optional" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The property describes a creating time of the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-migrationprojectcreationtime
	//
	// Deprecated: this property has been deprecated.
	MigrationProjectCreationTime *string `field:"optional" json:"migrationProjectCreationTime" yaml:"migrationProjectCreationTime"`
	// The property describes an identifier for the migration project.
	//
	// It is used for describing/deleting/modifying can be name/arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-migrationprojectidentifier
	//
	MigrationProjectIdentifier *string `field:"optional" json:"migrationProjectIdentifier" yaml:"migrationProjectIdentifier"`
	// The property describes a name to identify the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-migrationprojectname
	//
	MigrationProjectName *string `field:"optional" json:"migrationProjectName" yaml:"migrationProjectName"`
	// The property describes schema conversion application attributes for the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-schemaconversionapplicationattributes
	//
	SchemaConversionApplicationAttributes interface{} `field:"optional" json:"schemaConversionApplicationAttributes" yaml:"schemaConversionApplicationAttributes"`
	// The property describes source data provider descriptors for the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-sourcedataproviderdescriptors
	//
	SourceDataProviderDescriptors interface{} `field:"optional" json:"sourceDataProviderDescriptors" yaml:"sourceDataProviderDescriptors"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The property describes target data provider descriptors for the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-targetdataproviderdescriptors
	//
	TargetDataProviderDescriptors interface{} `field:"optional" json:"targetDataProviderDescriptors" yaml:"targetDataProviderDescriptors"`
	// The property describes transformation rules for the migration project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-migrationproject.html#cfn-dms-migrationproject-transformationrules
	//
	TransformationRules *string `field:"optional" json:"transformationRules" yaml:"transformationRules"`
}

