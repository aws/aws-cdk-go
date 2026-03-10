package previewawsbatchevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	ContainerOverrides: &ContainerOverrides{
//   		Environment: []ContainerOverridesItem{
//   			&ContainerOverridesItem{
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	JobDefinition: []*string{
//   		jsii.String("jobDefinition"),
//   	},
//   	JobName: []*string{
//   		jsii.String("jobName"),
//   	},
//   	JobQueue: []*string{
//   		jsii.String("jobQueue"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParameters struct {
	// containerOverrides property.
	//
	// Specify an array of string values to match this event if the actual value of containerOverrides is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerOverrides *AWSAPICallViaCloudTrail_ContainerOverrides `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// jobDefinition property.
	//
	// Specify an array of string values to match this event if the actual value of jobDefinition is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobDefinition *[]*string `field:"optional" json:"jobDefinition" yaml:"jobDefinition"`
	// jobName property.
	//
	// Specify an array of string values to match this event if the actual value of jobName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobName *[]*string `field:"optional" json:"jobName" yaml:"jobName"`
	// jobQueue property.
	//
	// Specify an array of string values to match this event if the actual value of jobQueue is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	JobQueue *[]*string `field:"optional" json:"jobQueue" yaml:"jobQueue"`
}

