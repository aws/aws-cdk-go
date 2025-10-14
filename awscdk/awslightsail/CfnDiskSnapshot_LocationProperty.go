package awslightsail


// The AWS Region and Availability Zone where the disk snapshot was created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &LocationProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	RegionName: jsii.String("regionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disksnapshot-location.html
//
type CfnDiskSnapshot_LocationProperty struct {
	// The Availability Zone where the disk snapshot was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disksnapshot-location.html#cfn-lightsail-disksnapshot-location-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AWS Region where the disk snapshot was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-disksnapshot-location.html#cfn-lightsail-disksnapshot-location-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

