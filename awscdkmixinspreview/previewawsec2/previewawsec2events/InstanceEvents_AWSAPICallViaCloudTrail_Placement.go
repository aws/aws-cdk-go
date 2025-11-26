package previewawsec2events


// Type definition for Placement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   placement := &Placement{
//   	AvailabilityZone: []*string{
//   		jsii.String("availabilityZone"),
//   	},
//   	Tenancy: []*string{
//   		jsii.String("tenancy"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_Placement struct {
	// availabilityZone property.
	//
	// Specify an array of string values to match this event if the actual value of availabilityZone is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZone *[]*string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// tenancy property.
	//
	// Specify an array of string values to match this event if the actual value of tenancy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tenancy *[]*string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

