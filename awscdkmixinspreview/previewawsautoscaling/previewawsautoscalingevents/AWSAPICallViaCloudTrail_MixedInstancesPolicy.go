package previewawsautoscalingevents


// Type definition for MixedInstancesPolicy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var overrides interface{}
//
//   mixedInstancesPolicy := &MixedInstancesPolicy{
//   	InstancesDistribution: &InstancesDistribution{
//   		OnDemandAllocationStrategy: []*string{
//   			jsii.String("onDemandAllocationStrategy"),
//   		},
//   		OnDemandBaseCapacity: []*string{
//   			jsii.String("onDemandBaseCapacity"),
//   		},
//   		OnDemandPercentageAboveBaseCapacity: []*string{
//   			jsii.String("onDemandPercentageAboveBaseCapacity"),
//   		},
//   		SpotAllocationStrategy: []*string{
//   			jsii.String("spotAllocationStrategy"),
//   		},
//   		SpotInstancePools: []*string{
//   			jsii.String("spotInstancePools"),
//   		},
//   	},
//   	LaunchTemplate: &LaunchTemplate1{
//   		LaunchTemplateSpecification: &LaunchTemplateSpecification{
//   			LaunchTemplateName: []*string{
//   				jsii.String("launchTemplateName"),
//   			},
//   			Version: []*string{
//   				jsii.String("version"),
//   			},
//   		},
//   		Overrides: []interface{}{
//   			overrides,
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_MixedInstancesPolicy struct {
	// instancesDistribution property.
	//
	// Specify an array of string values to match this event if the actual value of instancesDistribution is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstancesDistribution *AWSAPICallViaCloudTrail_InstancesDistribution `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
	// launchTemplate property.
	//
	// Specify an array of string values to match this event if the actual value of launchTemplate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplate *AWSAPICallViaCloudTrail_LaunchTemplate1 `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
}

