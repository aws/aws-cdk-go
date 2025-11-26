package previewawsec2events


// Type definition for CapacityReservationSpecification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   capacityReservationSpecification := &CapacityReservationSpecification{
//   	CapacityReservationPreference: []*string{
//   		jsii.String("capacityReservationPreference"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_CapacityReservationSpecification struct {
	// capacityReservationPreference property.
	//
	// Specify an array of string values to match this event if the actual value of capacityReservationPreference is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CapacityReservationPreference *[]*string `field:"optional" json:"capacityReservationPreference" yaml:"capacityReservationPreference"`
}

