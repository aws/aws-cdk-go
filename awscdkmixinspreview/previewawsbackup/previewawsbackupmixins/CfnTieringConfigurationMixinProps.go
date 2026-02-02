package previewawsbackupmixins


// Properties for CfnTieringConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTieringConfigurationMixinProps := &CfnTieringConfigurationMixinProps{
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
//   	TieringConfigurationTags: map[string]*string{
//   		"tieringConfigurationTagsKey": jsii.String("tieringConfigurationTags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html
//
type CfnTieringConfigurationMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html#cfn-backup-tieringconfiguration-backupvaultname
	//
	BackupVaultName *string `field:"optional" json:"backupVaultName" yaml:"backupVaultName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html#cfn-backup-tieringconfiguration-resourceselection
	//
	ResourceSelection interface{} `field:"optional" json:"resourceSelection" yaml:"resourceSelection"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html#cfn-backup-tieringconfiguration-tieringconfigurationname
	//
	TieringConfigurationName *string `field:"optional" json:"tieringConfigurationName" yaml:"tieringConfigurationName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-tieringconfiguration.html#cfn-backup-tieringconfiguration-tieringconfigurationtags
	//
	TieringConfigurationTags *map[string]*string `field:"optional" json:"tieringConfigurationTags" yaml:"tieringConfigurationTags"`
}

