package awsec2


// The maintenance options of your instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceOptionsProperty := &MaintenanceOptionsProperty{
//   	AutoRecovery: jsii.String("autoRecovery"),
//   	RebootMigration: jsii.String("rebootMigration"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-maintenanceoptions.html
//
type CfnLaunchTemplate_MaintenanceOptionsProperty struct {
	// Disables the automatic recovery behavior of your instance or sets it to default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-maintenanceoptions.html#cfn-ec2-launchtemplate-maintenanceoptions-autorecovery
	//
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-maintenanceoptions.html#cfn-ec2-launchtemplate-maintenanceoptions-rebootmigration
	//
	RebootMigration *string `field:"optional" json:"rebootMigration" yaml:"rebootMigration"`
}

