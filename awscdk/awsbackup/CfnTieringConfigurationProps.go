package awsbackup


// Properties for defining a `CfnTieringConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTieringConfigurationProps := &CfnTieringConfigurationProps{
//   	BackupVaultName: jsii.String("backupVaultName"),
//   	ResourceSelection: []interface{}{
//   		&ResourceSelectionProperty{
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   			ResourceType: jsii.String("resourceType"),
//   			TieringDownSettingsInDays: jsii.Number(123),
//   		},
//   	},
//   	TieringConfigurationName: jsii.String("tieringConfigurationName"),
//
//   	// the properties below are optional
//   	TieringConfigurationTags: map[string]*string{
//   		"tieringConfigurationTagsKey": jsii.String("tieringConfigurationTags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html
//
type CfnTieringConfigurationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html#cfn-backup-tieringconfiguration-backupvaultname
	//
	BackupVaultName *string `field:"required" json:"backupVaultName" yaml:"backupVaultName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html#cfn-backup-tieringconfiguration-resourceselection
	//
	ResourceSelection interface{} `field:"required" json:"resourceSelection" yaml:"resourceSelection"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html#cfn-backup-tieringconfiguration-tieringconfigurationname
	//
	TieringConfigurationName *string `field:"required" json:"tieringConfigurationName" yaml:"tieringConfigurationName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html#cfn-backup-tieringconfiguration-tieringconfigurationtags
	//
	TieringConfigurationTags *map[string]*string `field:"optional" json:"tieringConfigurationTags" yaml:"tieringConfigurationTags"`
}

