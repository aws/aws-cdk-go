package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCustomDBEngineVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomDBEngineVersionProps := &CfnCustomDBEngineVersionProps{
//   	DatabaseInstallationFilesS3BucketName: jsii.String("databaseInstallationFilesS3BucketName"),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//
//   	// the properties below are optional
//   	DatabaseInstallationFilesS3Prefix: jsii.String("databaseInstallationFilesS3Prefix"),
//   	Description: jsii.String("description"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Manifest: jsii.String("manifest"),
//   	Status: jsii.String("status"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html
//
type CfnCustomDBEngineVersionProps struct {
	// The name of an Amazon S3 bucket that contains database installation files for your CEV.
	//
	// For example, a valid bucket name is `my-custom-installation-files` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html#cfn-rds-customdbengineversion-databaseinstallationfiless3bucketname
	//
	DatabaseInstallationFilesS3BucketName *string `field:"required" json:"databaseInstallationFilesS3BucketName" yaml:"databaseInstallationFilesS3BucketName"`
	// The database engine to use for your custom engine version (CEV).
	//
	// Valid values:
	//
	// - `custom-oracle-ee`
	// - `custom-oracle-ee-cdb`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html#cfn-rds-customdbengineversion-engine
	//
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The name of your CEV.
	//
	// The name format is `major version.customized_string` . For example, a valid CEV name is `19.my_cev1` . This setting is required for RDS Custom for Oracle, but optional for Amazon RDS. The combination of `Engine` and `EngineVersion` is unique per customer per Region.
	//
	// *Constraints:* Minimum length is 1. Maximum length is 60.
	//
	// *Pattern:* `^[a-z0-9_.-]{1,60$` }
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html#cfn-rds-customdbengineversion-engineversion
	//
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// The Amazon S3 directory that contains the database installation files for your CEV.
	//
	// For example, a valid bucket name is `123456789012/cev1` . If this setting isn't specified, no prefix is assumed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html#cfn-rds-customdbengineversion-databaseinstallationfiless3prefix
	//
	DatabaseInstallationFilesS3Prefix *string `field:"optional" json:"databaseInstallationFilesS3Prefix" yaml:"databaseInstallationFilesS3Prefix"`
	// An optional description of your CEV.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html#cfn-rds-customdbengineversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS KMS key identifier for an encrypted CEV.
	//
	// A symmetric encryption KMS key is required for RDS Custom, but optional for Amazon RDS.
	//
	// If you have an existing symmetric encryption KMS key in your account, you can use it with RDS Custom. No further action is necessary. If you don't already have a symmetric encryption KMS key in your account, follow the instructions in [Creating a symmetric encryption KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html#create-symmetric-cmk) in the *AWS Key Management Service Developer Guide* .
	//
	// You can choose the same symmetric encryption key when you create a CEV and a DB instance, or choose different keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html#cfn-rds-customdbengineversion-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3. Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which they are listed.
	//
	// The following JSON fields are valid:
	//
	// - **MediaImportTemplateVersion** - Version of the CEV manifest. The date is in the format `YYYY-MM-DD` .
	// - **databaseInstallationFileNames** - Ordered list of installation files for the CEV.
	// - **opatchFileNames** - Ordered list of OPatch installers used for the Oracle DB engine.
	// - **psuRuPatchFileNames** - The PSU and RU patches for this CEV.
	// - **OtherPatchFileNames** - The patches that are not in the list of PSU and RU patches. Amazon RDS applies these patches after applying the PSU and RU patches.
	//
	// For more information, see [Creating the CEV manifest](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-cev.html#custom-cev.preparing.manifest) in the *Amazon RDS User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html#cfn-rds-customdbengineversion-manifest
	//
	Manifest *string `field:"optional" json:"manifest" yaml:"manifest"`
	// A value that indicates the status of a custom engine version (CEV).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html#cfn-rds-customdbengineversion-status
	//
	// Default: - "available".
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of tags.
	//
	// For more information, see [Tagging Amazon RDS Resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-customdbengineversion.html#cfn-rds-customdbengineversion-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

