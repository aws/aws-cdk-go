package awsbackup


// Specifies an object containing an array of `Transition` objects that determine how long in days before a recovery point transitions to cold storage or is deleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleResourceTypeProperty := &lifecycleResourceTypeProperty{
//   	deleteAfterDays: jsii.Number(123),
//   	moveToColdStorageAfterDays: jsii.Number(123),
//   }
//
type CfnBackupPlan_LifecycleResourceTypeProperty struct {
	// Specifies the number of days after creation that a recovery point is deleted.
	//
	// Must be greater than `MoveToColdStorageAfterDays` .
	DeleteAfterDays *float64 `field:"optional" json:"deleteAfterDays" yaml:"deleteAfterDays"`
	// Specifies the number of days after creation that a recovery point is moved to cold storage.
	MoveToColdStorageAfterDays *float64 `field:"optional" json:"moveToColdStorageAfterDays" yaml:"moveToColdStorageAfterDays"`
}

