package previewawselasticloadbalancingevents


// Type definition for RequestParametersItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParametersItem := &RequestParametersItem{
//   	InstancePort: []*string{
//   		jsii.String("instancePort"),
//   	},
//   	LoadBalancerPort: []*string{
//   		jsii.String("loadBalancerPort"),
//   	},
//   	Protocol: []*string{
//   		jsii.String("protocol"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParametersItem struct {
	// instancePort property.
	//
	// Specify an array of string values to match this event if the actual value of instancePort is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstancePort *[]*string `field:"optional" json:"instancePort" yaml:"instancePort"`
	// loadBalancerPort property.
	//
	// Specify an array of string values to match this event if the actual value of loadBalancerPort is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LoadBalancerPort *[]*string `field:"optional" json:"loadBalancerPort" yaml:"loadBalancerPort"`
	// protocol property.
	//
	// Specify an array of string values to match this event if the actual value of protocol is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Protocol *[]*string `field:"optional" json:"protocol" yaml:"protocol"`
}

