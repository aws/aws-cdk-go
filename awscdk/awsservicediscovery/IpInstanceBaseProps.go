package awsservicediscovery


// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewHttpNamespace(stack, jsii.String("MyNamespace"), &HttpNamespaceProps{
//   	Name: jsii.String("MyHTTPNamespace"),
//   })
//
//   service1 := namespace.CreateService(jsii.String("NonIpService"), &BaseServiceProps{
//   	Description: jsii.String("service registering non-ip instances"),
//   })
//
//   service1.RegisterNonIpInstance(jsii.String("NonIpInstance"), &NonIpInstanceBaseProps{
//   	CustomAttributes: map[string]*string{
//   		"arn": jsii.String("arn:aws:s3:::mybucket"),
//   	},
//   })
//
//   service2 := namespace.CreateService(jsii.String("IpService"), &BaseServiceProps{
//   	Description: jsii.String("service registering ip instances"),
//   	HealthCheck: &HealthCheckConfig{
//   		Type: servicediscovery.HealthCheckType_HTTP,
//   		ResourcePath: jsii.String("/check"),
//   	},
//   })
//
//   service2.RegisterIpInstance(jsii.String("IpInstance"), &IpInstanceBaseProps{
//   	Ipv4: jsii.String("54.239.25.192"),
//   })
//
//   app.Synth()
//
type IpInstanceBaseProps struct {
	// Custom attributes of the instance.
	// Default: none.
	//
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	// Default: Automatically generated name.
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// If the service that you specify contains a template for an A record, the IPv4 address that you want AWS Cloud Map to use for the value of the A record.
	// Default: none.
	//
	Ipv4 *string `field:"optional" json:"ipv4" yaml:"ipv4"`
	// If the service that you specify contains a template for an AAAA record, the IPv6 address that you want AWS Cloud Map to use for the value of the AAAA record.
	// Default: none.
	//
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// The port on the endpoint that you want AWS Cloud Map to perform health checks on.
	//
	// This value is also used for
	// the port value in an SRV record if the service that you specify includes an SRV record. You can also specify a
	// default port that is applied to all instances in the Service configuration.
	// Default: 80.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

