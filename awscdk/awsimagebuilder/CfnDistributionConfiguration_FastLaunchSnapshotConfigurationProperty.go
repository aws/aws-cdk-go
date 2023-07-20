package awsimagebuilder


// Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-fastlaunchsnapshotconfiguration.html
//
type CfnDistributionConfiguration_FastLaunchSnapshotConfigurationProperty struct {
	// The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled Windows AMI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-fastlaunchsnapshotconfiguration.html#cfn-imagebuilder-distributionconfiguration-fastlaunchsnapshotconfiguration-targetresourcecount
	//
	TargetResourceCount *float64 `field:"optional" json:"targetResourceCount" yaml:"targetResourceCount"`
}

