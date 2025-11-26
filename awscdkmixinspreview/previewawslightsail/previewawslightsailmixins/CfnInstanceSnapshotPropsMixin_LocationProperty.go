package previewawslightsailmixins


// The region name and Availability Zone where you created the snapshot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   locationProperty := &LocationProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	RegionName: jsii.String("regionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instancesnapshot-location.html
//
type CfnInstanceSnapshotPropsMixin_LocationProperty struct {
	// The Availability Zone.
	//
	// Follows the format us-east-2a (case-sensitive).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instancesnapshot-location.html#cfn-lightsail-instancesnapshot-location-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AWS Region name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-instancesnapshot-location.html#cfn-lightsail-instancesnapshot-location-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

