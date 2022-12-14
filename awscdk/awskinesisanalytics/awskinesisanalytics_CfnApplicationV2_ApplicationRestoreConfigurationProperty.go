package awskinesisanalytics


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
type CfnApplicationV2_ApplicationRestoreConfigurationProperty struct {
	// `CfnApplicationV2.ApplicationRestoreConfigurationProperty.ApplicationRestoreType`.
	ApplicationRestoreType *string `field:"required" json:"applicationRestoreType" yaml:"applicationRestoreType"`
	// `CfnApplicationV2.ApplicationRestoreConfigurationProperty.SnapshotName`.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
}

