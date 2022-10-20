package awsservicediscovery


// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import servicediscovery "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewHttpNamespace(stack, jsii.String("MyNamespace"), &httpNamespaceProps{
//   	name: jsii.String("MyHTTPNamespace"),
//   })
//
//   service1 := namespace.createService(jsii.String("NonIpService"), &baseServiceProps{
//   	description: jsii.String("service registering non-ip instances"),
//   })
//
//   service1.registerNonIpInstance(jsii.String("NonIpInstance"), &nonIpInstanceBaseProps{
//   	customAttributes: map[string]*string{
//   		"arn": jsii.String("arn:aws:s3:::mybucket"),
//   	},
//   })
//
//   service2 := namespace.createService(jsii.String("IpService"), &baseServiceProps{
//   	description: jsii.String("service registering ip instances"),
//   	healthCheck: &healthCheckConfig{
//   		type: servicediscovery.healthCheckType_HTTP,
//   		resourcePath: jsii.String("/check"),
//   	},
//   })
//
//   service2.registerIpInstance(jsii.String("IpInstance"), &ipInstanceBaseProps{
//   	ipv4: jsii.String("54.239.25.192"),
//   })
//
//   app.synth()
//
type IpInstanceBaseProps struct {
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
}

