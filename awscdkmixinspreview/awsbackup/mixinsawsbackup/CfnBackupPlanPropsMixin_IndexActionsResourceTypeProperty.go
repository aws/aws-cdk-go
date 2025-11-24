package mixinsawsbackup


// Specifies index actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   indexActionsResourceTypeProperty := &IndexActionsResourceTypeProperty{
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-indexactionsresourcetype.html
//
type CfnBackupPlanPropsMixin_IndexActionsResourceTypeProperty struct {
	// 0 or 1 index action will be accepted for each BackupRule.
	//
	// Valid values:
	//
	// - `EBS` for Amazon Elastic Block Store
	// - `S3` for Amazon Simple Storage Service (Amazon S3).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-indexactionsresourcetype.html#cfn-backup-backupplan-indexactionsresourcetype-resourcetypes
	//
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

