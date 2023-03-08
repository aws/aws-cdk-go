package awswaf


// Properties for defining a `CfnIPSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPSetProps := &cfnIPSetProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ipSetDescriptors: []interface{}{
//   		map[string]*string{
//   			"type": jsii.String("type"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnIPSetProps struct {
	// The name of the `IPSet` .
	//
	// You can't change the name of an `IPSet` after you create it.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The IP address type ( `IPV4` or `IPV6` ) and the IP address range (in CIDR notation) that web requests originate from.
	//
	// If the `WebACL` is associated with an Amazon CloudFront distribution and the viewer did not use an HTTP proxy or a load balancer to send the request, this is the value of the c-ip field in the CloudFront access logs.
	IpSetDescriptors interface{} `field:"optional" json:"ipSetDescriptors" yaml:"ipSetDescriptors"`
}

