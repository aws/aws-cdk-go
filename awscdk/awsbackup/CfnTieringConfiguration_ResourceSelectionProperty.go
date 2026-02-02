package awsbackup


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnTieringConfiguration_ResourceSelectionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-tieringconfiguration-resourceselection.html#cfn-backup-tieringconfiguration-resourceselection-resources
	//
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-tieringconfiguration-resourceselection.html#cfn-backup-tieringconfiguration-resourceselection-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-tieringconfiguration-resourceselection.html#cfn-backup-tieringconfiguration-resourceselection-tieringdownsettingsindays
	//
	TieringDownSettingsInDays *float64 `field:"required" json:"tieringDownSettingsInDays" yaml:"tieringDownSettingsInDays"`
}

