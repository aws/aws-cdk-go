package awsservicediscovery


// Basic props needed to create a service in a given namespace.
//
// Used by HttpNamespace.createService
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
type BaseServiceProps struct {
	// Structure containing failure threshold for a custom health checker.
	//
	// Only one of healthCheckConfig or healthCheckCustomConfig can be specified.
	// See: https://docs.aws.amazon.com/cloud-map/latest/api/API_HealthCheckCustomConfig.html
	CustomHealthCheck *HealthCheckCustomConfig `field:"optional" json:"customHealthCheck" yaml:"customHealthCheck"`
	// A description of the service.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Settings for an optional health check.
	//
	// If you specify health check settings, AWS Cloud Map associates the health
	// check with the records that you specify in DnsConfig. Only one of healthCheckConfig or healthCheckCustomConfig can
	// be specified. Not valid for PrivateDnsNamespaces. If you use healthCheck, you can only register IP instances to
	// this service.
	HealthCheck *HealthCheckConfig `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// A name for the Service.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

