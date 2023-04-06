package awsquicksight


// The drill down options for data points in a dashbaord.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPointDrillUpDownOptionProperty := &DataPointDrillUpDownOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_DataPointDrillUpDownOptionProperty struct {
	// The status of the drill down options of data points.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

