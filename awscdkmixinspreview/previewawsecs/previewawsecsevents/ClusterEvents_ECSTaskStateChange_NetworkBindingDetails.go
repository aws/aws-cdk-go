package previewawsecsevents


// Type definition for NetworkBindingDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkBindingDetails := &NetworkBindingDetails{
//   	BindIp: []*string{
//   		jsii.String("bindIp"),
//   	},
//   	ContainerPort: []*string{
//   		jsii.String("containerPort"),
//   	},
//   	HostPort: []*string{
//   		jsii.String("hostPort"),
//   	},
//   	Protocol: []*string{
//   		jsii.String("protocol"),
//   	},
//   }
//
// Experimental.
type ClusterEvents_ECSTaskStateChange_NetworkBindingDetails struct {
	// bindIP property.
	//
	// Specify an array of string values to match this event if the actual value of bindIP is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BindIp *[]*string `field:"optional" json:"bindIp" yaml:"bindIp"`
	// containerPort property.
	//
	// Specify an array of string values to match this event if the actual value of containerPort is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ContainerPort *[]*string `field:"optional" json:"containerPort" yaml:"containerPort"`
	// hostPort property.
	//
	// Specify an array of string values to match this event if the actual value of hostPort is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HostPort *[]*string `field:"optional" json:"hostPort" yaml:"hostPort"`
	// protocol property.
	//
	// Specify an array of string values to match this event if the actual value of protocol is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Protocol *[]*string `field:"optional" json:"protocol" yaml:"protocol"`
}

