package previewawsbackupmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceSelectionProperty := &ResourceSelectionProperty{
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	TieringDownSettingsInDays: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-tieringconfiguration-resourceselection.html
//
type CfnTieringConfigurationPropsMixin_ResourceSelectionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-tieringconfiguration-resourceselection.html#cfn-backup-tieringconfiguration-resourceselection-resources
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-tieringconfiguration-resourceselection.html#cfn-backup-tieringconfiguration-resourceselection-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-tieringconfiguration-resourceselection.html#cfn-backup-tieringconfiguration-resourceselection-tieringdownsettingsindays
	//
	TieringDownSettingsInDays *float64 `field:"optional" json:"tieringDownSettingsInDays" yaml:"tieringDownSettingsInDays"`
}

