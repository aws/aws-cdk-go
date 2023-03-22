package awsec2


// The maintenance options of your instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   maintenanceOptionsProperty := &maintenanceOptionsProperty{
//   	autoRecovery: jsii.String("autoRecovery"),
//   }
//
type CfnLaunchTemplate_MaintenanceOptionsProperty struct {
	// Disables the automatic recovery behavior of your instance or sets it to default.
	AutoRecovery *string `field:"optional" json:"autoRecovery" yaml:"autoRecovery"`
}

