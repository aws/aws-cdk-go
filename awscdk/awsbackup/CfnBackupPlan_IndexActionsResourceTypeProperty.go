package awsbackup


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   indexActionsResourceTypeProperty := &IndexActionsResourceTypeProperty{
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-indexactionsresourcetype.html
//
type CfnBackupPlan_IndexActionsResourceTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-indexactionsresourcetype.html#cfn-backup-backupplan-indexactionsresourcetype-resourcetypes
	//
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

