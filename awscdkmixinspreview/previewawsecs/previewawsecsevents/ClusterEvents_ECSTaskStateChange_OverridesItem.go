package previewawsecsevents


// Type definition for OverridesItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   overridesItem := &OverridesItem{
//   	Command: []*string{
//   		jsii.String("command"),
//   	},
//   	Cpu: []*string{
//   		jsii.String("cpu"),
//   	},
//   	Environment: []map[string]*string{
//   		map[string]*string{
//   			"environmentKey": jsii.String("environment"),
//   		},
//   	},
//   	Memory: []*string{
//   		jsii.String("memory"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSTaskStateChange_OverridesItem struct {
	// command property.
	//
	// Specify an array of string values to match this event if the actual value of command is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// cpu property.
	//
	// Specify an array of string values to match this event if the actual value of cpu is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Cpu *[]*string `field:"optional" json:"cpu" yaml:"cpu"`
	// environment property.
	//
	// Specify an array of string values to match this event if the actual value of environment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Environment *[]*map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// memory property.
	//
	// Specify an array of string values to match this event if the actual value of memory is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Memory *[]*string `field:"optional" json:"memory" yaml:"memory"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
}

