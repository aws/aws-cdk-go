package previewawsec2events


// Type definition for TargetCapacitySpecification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetCapacitySpecification := &TargetCapacitySpecification{
//   	DefaultTargetCapacityType: []*string{
//   		jsii.String("defaultTargetCapacityType"),
//   	},
//   	OnDemandTargetCapacity: []*string{
//   		jsii.String("onDemandTargetCapacity"),
//   	},
//   	SpotTargetCapacity: []*string{
//   		jsii.String("spotTargetCapacity"),
//   	},
//   	TotalTargetCapacity: []*string{
//   		jsii.String("totalTargetCapacity"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_TargetCapacitySpecification struct {
	// DefaultTargetCapacityType property.
	//
	// Specify an array of string values to match this event if the actual value of DefaultTargetCapacityType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DefaultTargetCapacityType *[]*string `field:"optional" json:"defaultTargetCapacityType" yaml:"defaultTargetCapacityType"`
	// OnDemandTargetCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of OnDemandTargetCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OnDemandTargetCapacity *[]*string `field:"optional" json:"onDemandTargetCapacity" yaml:"onDemandTargetCapacity"`
	// SpotTargetCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of SpotTargetCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotTargetCapacity *[]*string `field:"optional" json:"spotTargetCapacity" yaml:"spotTargetCapacity"`
	// TotalTargetCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of TotalTargetCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TotalTargetCapacity *[]*string `field:"optional" json:"totalTargetCapacity" yaml:"totalTargetCapacity"`
}

