package previewawslightsailmixins


// The AWS Region and Availability Zone where the disk is located.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-location.html
//
type CfnDiskPropsMixin_LocationProperty struct {
	// The Availability Zone where the disk is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-location.html#cfn-lightsail-disk-location-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AWS Region where the disk is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disk-location.html#cfn-lightsail-disk-location-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

