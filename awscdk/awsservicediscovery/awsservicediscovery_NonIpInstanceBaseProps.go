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
type NonIpInstanceBaseProps struct {
	// Custom attributes of the instance.
	CustomAttributes *map[string]*string `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// The id of the instance resource.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
}

