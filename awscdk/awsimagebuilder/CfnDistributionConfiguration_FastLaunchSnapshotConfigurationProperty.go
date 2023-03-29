package awsimagebuilder


// Configuration settings for creating and managing pre-provisioned snapshots for a fast-launch enabled Windows AMI.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fastLaunchSnapshotConfigurationProperty := &FastLaunchSnapshotConfigurationProperty{
//   	TargetResourceCount: jsii.Number(123),
//   }
//
type CfnDistributionConfiguration_FastLaunchSnapshotConfigurationProperty struct {
	// The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.
	TargetResourceCount *float64 `field:"optional" json:"targetResourceCount" yaml:"targetResourceCount"`
}

