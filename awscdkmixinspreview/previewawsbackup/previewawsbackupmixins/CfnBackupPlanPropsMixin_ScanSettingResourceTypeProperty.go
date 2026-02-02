package previewawsbackupmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scanSettingResourceTypeProperty := &ScanSettingResourceTypeProperty{
//   	MalwareScanner: jsii.String("malwareScanner"),
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	ScannerRoleArn: jsii.String("scannerRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-scansettingresourcetype.html
//
type CfnBackupPlanPropsMixin_ScanSettingResourceTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-scansettingresourcetype.html#cfn-backup-backupplan-scansettingresourcetype-malwarescanner
	//
	MalwareScanner *string `field:"optional" json:"malwareScanner" yaml:"malwareScanner"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-scansettingresourcetype.html#cfn-backup-backupplan-scansettingresourcetype-resourcetypes
	//
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-scansettingresourcetype.html#cfn-backup-backupplan-scansettingresourcetype-scannerrolearn
	//
	ScannerRoleArn *string `field:"optional" json:"scannerRoleArn" yaml:"scannerRoleArn"`
}

