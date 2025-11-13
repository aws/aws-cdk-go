package interfacesawsbackup


// A reference to a BackupPlan resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupPlanReference := &BackupPlanReference{
//   	BackupPlanArn: jsii.String("backupPlanArn"),
//   	BackupPlanId: jsii.String("backupPlanId"),
//   }
//
type BackupPlanReference struct {
	// The ARN of the BackupPlan resource.
	BackupPlanArn *string `field:"required" json:"backupPlanArn" yaml:"backupPlanArn"`
	// The BackupPlanId of the BackupPlan resource.
	BackupPlanId *string `field:"required" json:"backupPlanId" yaml:"backupPlanId"`
}

