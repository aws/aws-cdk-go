package previewawsautoscalingevents


// Type definition for Details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   details := &Details{
//   	AvailabilityZone: []*string{
//   		jsii.String("availabilityZone"),
//   	},
//   	SubnetId: []*string{
//   		jsii.String("subnetId"),
//   	},
//   }
//
// Experimental.
type AutoScalingGroupEvents_EC2InstanceLaunchSuccessful_Details struct {
	// Availability Zone property.
	//
	// Specify an array of string values to match this event if the actual value of Availability Zone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZone *[]*string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Subnet ID property.
	//
	// Specify an array of string values to match this event if the actual value of Subnet ID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SubnetId *[]*string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

