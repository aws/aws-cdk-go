package awselasticloadbalancingv2


// Information about a host header condition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostHeaderConfigProperty := &hostHeaderConfigProperty{
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_HostHeaderConfigProperty struct {
	// One or more host names.
	//
	// The maximum size of each name is 128 characters. The comparison is case insensitive. The following wildcard characters are supported: * (matches 0 or more characters) and ? (matches exactly 1 character).
	//
	// If you specify multiple strings, the condition is satisfied if one of the strings matches the host name.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

