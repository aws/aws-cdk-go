package awsappstream


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumeConfigProperty := &VolumeConfigProperty{
//   	VolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-volumeconfig.html
//
type CfnFleet_VolumeConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-volumeconfig.html#cfn-appstream-fleet-volumeconfig-volumesizeingb
	//
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

