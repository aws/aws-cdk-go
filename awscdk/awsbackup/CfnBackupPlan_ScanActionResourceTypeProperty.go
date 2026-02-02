package awsbackup


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scanActionResourceTypeProperty := &ScanActionResourceTypeProperty{
//   	MalwareScanner: jsii.String("malwareScanner"),
//   	ScanMode: jsii.String("scanMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-scanactionresourcetype.html
//
type CfnBackupPlan_ScanActionResourceTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-scanactionresourcetype.html#cfn-backup-backupplan-scanactionresourcetype-malwarescanner
	//
	MalwareScanner *string `field:"optional" json:"malwareScanner" yaml:"malwareScanner"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-scanactionresourcetype.html#cfn-backup-backupplan-scanactionresourcetype-scanmode
	//
	ScanMode *string `field:"optional" json:"scanMode" yaml:"scanMode"`
}

