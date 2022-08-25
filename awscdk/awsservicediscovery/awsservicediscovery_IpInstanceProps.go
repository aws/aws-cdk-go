package awsservicediscovery


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var service service
//
//   ipInstanceProps := &ipInstanceProps{
//   	service: service,
//
//   	// the properties below are optional
//   	customAttributes: map[string]*string{
//   		"customAttributesKey": jsii.String("customAttributes"),
//   	},
//   	instanceId: jsii.String("instanceId"),
//   	ipv4: jsii.String("ipv4"),
//   	ipv6: jsii.String("ipv6"),
//   	port: jsii.Number(123),
//   }
//
type IpInstanceProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// If the service that you specify contains a template for an A record, the IPv4 address that you want AWS Cloud Map to use for the value of the A record.
	Ipv4 *string `field:"optional" json:"ipv4" yaml:"ipv4"`
	// If the service that you specify contains a template for an AAAA record, the IPv6 address that you want AWS Cloud Map to use for the value of the AAAA record.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// The port on the endpoint that you want AWS Cloud Map to perform health checks on.
	//
	// This value is also used for
	// the port value in an SRV record if the service that you specify includes an SRV record. You can also specify a
	// default port that is applied to all instances in the Service configuration.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The Cloudmap service this resource is registered to.
	Service IService `field:"required" json:"service" yaml:"service"`
}

