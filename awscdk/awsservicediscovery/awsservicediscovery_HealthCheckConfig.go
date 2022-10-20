package awsservicediscovery


// Settings for an optional Amazon Route 53 health check.
//
// If you specify settings for a health check, AWS Cloud Map
// associates the health check with all the records that you specify in DnsConfig. Only valid with a PublicDnsNamespace.
//
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
type HealthCheckConfig struct {
	// The number of consecutive health checks that an endpoint must pass or fail for Route 53 to change the current status of the endpoint from unhealthy to healthy or vice versa.
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// The path that you want Route 53 to request when performing health checks.
	//
	// Do not use when health check type is TCP.
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// The type of health check that you want to create, which indicates how Route 53 determines whether an endpoint is healthy.
	//
	// Cannot be modified once created. Supported values are HTTP, HTTPS, and TCP.
	Type HealthCheckType `field:"optional" json:"type" yaml:"type"`
}

