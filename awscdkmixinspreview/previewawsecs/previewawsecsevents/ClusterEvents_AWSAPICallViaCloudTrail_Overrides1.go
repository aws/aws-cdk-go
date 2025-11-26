package previewawsecsevents


// Type definition for Overrides_1.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var environment interface{}
//   var inferenceAcceleratorOverrides interface{}
//   var resourceRequirements interface{}
//
//   overrides1 := &Overrides1{
//   	ContainerOverrides: []Overrides1Item{
//   		&Overrides1Item{
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: []*string{
//   				jsii.String("cpu"),
//   			},
//   			Environment: []interface{}{
//   				environment,
//   			},
//   			Memory: []*string{
//   				jsii.String("memory"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			ResourceRequirements: []interface{}{
//   				resourceRequirements,
//   			},
//   		},
//   	},
//   	InferenceAcceleratorOverrides: []interface{}{
//   		inferenceAcceleratorOverrides,
//   	},
//   }
//
// Experimental.
type ClusterEvents_AWSAPICallViaCloudTrail_Overrides1 struct {
	// containerOverrides property.
	//
	// Specify an array of string values to match this event if the actual value of containerOverrides is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerOverrides *[]*ClusterEvents_AWSAPICallViaCloudTrail_Overrides1Item `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// inferenceAcceleratorOverrides property.
	//
	// Specify an array of string values to match this event if the actual value of inferenceAcceleratorOverrides is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InferenceAcceleratorOverrides *[]interface{} `field:"optional" json:"inferenceAcceleratorOverrides" yaml:"inferenceAcceleratorOverrides"`
}

