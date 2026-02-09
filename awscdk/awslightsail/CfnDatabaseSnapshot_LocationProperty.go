package awslightsail


// The Region name and Availability Zone where the database snapshot is located.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-databasesnapshot-location.html
//
type CfnDatabaseSnapshot_LocationProperty struct {
	// The Availability Zone where the database snapshot was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-databasesnapshot-location.html#cfn-lightsail-databasesnapshot-location-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AWS Region where the database snapshot was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-databasesnapshot-location.html#cfn-lightsail-databasesnapshot-location-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}

