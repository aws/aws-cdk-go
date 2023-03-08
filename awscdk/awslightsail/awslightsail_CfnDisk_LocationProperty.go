package awslightsail


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &locationProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	regionName: jsii.String("regionName"),
//   }
//
type CfnDisk_LocationProperty struct {
	// `CfnDisk.LocationProperty.AvailabilityZone`.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// `CfnDisk.LocationProperty.RegionName`.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

