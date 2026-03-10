package previewawsbatchevents


// Type definition for ContainerOverrides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerOverrides := &ContainerOverrides{
//   	Environment: []ContainerOverridesItem{
//   		&ContainerOverridesItem{
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ContainerOverrides struct {
	// environment property.
	//
	// Specify an array of string values to match this event if the actual value of environment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Environment *[]*AWSAPICallViaCloudTrail_ContainerOverridesItem `field:"optional" json:"environment" yaml:"environment"`
}

