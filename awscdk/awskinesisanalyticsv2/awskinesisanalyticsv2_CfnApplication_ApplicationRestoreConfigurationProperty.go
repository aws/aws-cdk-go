package awskinesisanalyticsv2


// Specifies the method and snapshot to use when restarting an application using previously saved application state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationRestoreConfigurationProperty := &applicationRestoreConfigurationProperty{
//   	applicationRestoreType: jsii.String("applicationRestoreType"),
//
//   	// the properties below are optional
//   	snapshotName: jsii.String("snapshotName"),
//   }
//
type CfnApplication_ApplicationRestoreConfigurationProperty struct {
	// Specifies how the application should be restored.
	ApplicationRestoreType *string `field:"required" json:"applicationRestoreType" yaml:"applicationRestoreType"`
	// The identifier of an existing snapshot of application state to use to restart an application.
	//
	// The application uses this value if `RESTORE_FROM_CUSTOM_SNAPSHOT` is specified for the `ApplicationRestoreType` .
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
}

