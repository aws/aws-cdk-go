package previewawsecsevents


// Type definition for Overrides.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   overrides := &Overrides{
//   	ContainerOverrides: []OverridesItem{
//   		&OverridesItem{
//   			Command: []*string{
//   				jsii.String("command"),
//   			},
//   			Cpu: []*string{
//   				jsii.String("cpu"),
//   			},
//   			Environment: []map[string]*string{
//   				map[string]*string{
//   					"environmentKey": jsii.String("environment"),
//   				},
//   			},
//   			Memory: []*string{
//   				jsii.String("memory"),
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSTaskStateChange_Overrides struct {
	// containerOverrides property.
	//
	// Specify an array of string values to match this event if the actual value of containerOverrides is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerOverrides *[]*ClusterEvents_ECSTaskStateChange_OverridesItem `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
}

