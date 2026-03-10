package previewawsautoscalingevents


// Type definition for InstancesDistribution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instancesDistribution := &InstancesDistribution{
//   	OnDemandAllocationStrategy: []*string{
//   		jsii.String("onDemandAllocationStrategy"),
//   	},
//   	OnDemandBaseCapacity: []*string{
//   		jsii.String("onDemandBaseCapacity"),
//   	},
//   	OnDemandPercentageAboveBaseCapacity: []*string{
//   		jsii.String("onDemandPercentageAboveBaseCapacity"),
//   	},
//   	SpotAllocationStrategy: []*string{
//   		jsii.String("spotAllocationStrategy"),
//   	},
//   	SpotInstancePools: []*string{
//   		jsii.String("spotInstancePools"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_InstancesDistribution struct {
	// onDemandAllocationStrategy property.
	//
	// Specify an array of string values to match this event if the actual value of onDemandAllocationStrategy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OnDemandAllocationStrategy *[]*string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// onDemandBaseCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of onDemandBaseCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OnDemandBaseCapacity *[]*string `field:"optional" json:"onDemandBaseCapacity" yaml:"onDemandBaseCapacity"`
	// onDemandPercentageAboveBaseCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of onDemandPercentageAboveBaseCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	OnDemandPercentageAboveBaseCapacity *[]*string `field:"optional" json:"onDemandPercentageAboveBaseCapacity" yaml:"onDemandPercentageAboveBaseCapacity"`
	// spotAllocationStrategy property.
	//
	// Specify an array of string values to match this event if the actual value of spotAllocationStrategy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotAllocationStrategy *[]*string `field:"optional" json:"spotAllocationStrategy" yaml:"spotAllocationStrategy"`
	// spotInstancePools property.
	//
	// Specify an array of string values to match this event if the actual value of spotInstancePools is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotInstancePools *[]*string `field:"optional" json:"spotInstancePools" yaml:"spotInstancePools"`
}

