package mixinsawslightsail


// `AutoSnapshotAddOn` is a property of the [AddOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-addon.html) property. It describes the automatic snapshot add-on for a disk.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoSnapshotAddOnProperty := &AutoSnapshotAddOnProperty{
//   	SnapshotTimeOfDay: jsii.String("snapshotTimeOfDay"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-autosnapshotaddon.html
//
type CfnDiskPropsMixin_AutoSnapshotAddOnProperty struct {
	// The daily time when an automatic snapshot will be created.
	//
	// Constraints:
	//
	// - Must be in `HH:00` format, and in an hourly increment.
	// - Specified in Coordinated Universal Time (UTC).
	// - The snapshot will be automatically created between the time specified and up to 45 minutes after.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-autosnapshotaddon.html#cfn-lightsail-disk-autosnapshotaddon-snapshottimeofday
	//
	SnapshotTimeOfDay *string `field:"optional" json:"snapshotTimeOfDay" yaml:"snapshotTimeOfDay"`
}

