package previewawsec2events


// Type definition for CreateFleetResponse.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   createFleetResponse := &CreateFleetResponse{
//   	ErrorSet: []*string{
//   		jsii.String("errorSet"),
//   	},
//   	FleetId: []*string{
//   		jsii.String("fleetId"),
//   	},
//   	FleetInstanceSet: []*string{
//   		jsii.String("fleetInstanceSet"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	Xmlns: []*string{
//   		jsii.String("xmlns"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_CreateFleetResponse struct {
	// errorSet property.
	//
	// Specify an array of string values to match this event if the actual value of errorSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorSet *[]*string `field:"optional" json:"errorSet" yaml:"errorSet"`
	// fleetId property.
	//
	// Specify an array of string values to match this event if the actual value of fleetId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FleetId *[]*string `field:"optional" json:"fleetId" yaml:"fleetId"`
	// fleetInstanceSet property.
	//
	// Specify an array of string values to match this event if the actual value of fleetInstanceSet is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FleetInstanceSet *[]*string `field:"optional" json:"fleetInstanceSet" yaml:"fleetInstanceSet"`
	// requestId property.
	//
	// Specify an array of string values to match this event if the actual value of requestId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// xmlns property.
	//
	// Specify an array of string values to match this event if the actual value of xmlns is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Xmlns *[]*string `field:"optional" json:"xmlns" yaml:"xmlns"`
}

