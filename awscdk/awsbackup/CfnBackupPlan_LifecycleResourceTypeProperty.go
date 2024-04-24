package awsbackup


// Specifies an object containing an array of `Transition` objects that determine how long in days before a recovery point transitions to cold storage or is deleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleResourceTypeProperty := &LifecycleResourceTypeProperty{
//   	DeleteAfterDays: jsii.Number(123),
//   	MoveToColdStorageAfterDays: jsii.Number(123),
//   	OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-lifecycleresourcetype.html
//
type CfnBackupPlan_LifecycleResourceTypeProperty struct {
	// Specifies the number of days after creation that a recovery point is deleted.
	//
	// Must be greater than `MoveToColdStorageAfterDays` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-lifecycleresourcetype.html#cfn-backup-backupplan-lifecycleresourcetype-deleteafterdays
	//
	DeleteAfterDays *float64 `field:"optional" json:"deleteAfterDays" yaml:"deleteAfterDays"`
	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-lifecycleresourcetype.html#cfn-backup-backupplan-lifecycleresourcetype-movetocoldstorageafterdays
	//
	MoveToColdStorageAfterDays *float64 `field:"optional" json:"moveToColdStorageAfterDays" yaml:"moveToColdStorageAfterDays"`
	// If the value is true, your backup plan transitions supported resources to archive (cold) storage tier in accordance with your lifecycle settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-lifecycleresourcetype.html#cfn-backup-backupplan-lifecycleresourcetype-optintoarchiveforsupportedresources
	//
	OptInToArchiveForSupportedResources interface{} `field:"optional" json:"optInToArchiveForSupportedResources" yaml:"optInToArchiveForSupportedResources"`
}

