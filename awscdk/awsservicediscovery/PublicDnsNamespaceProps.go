package awsservicediscovery


// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("aws-servicediscovery-integ"))
//
//   namespace := servicediscovery.NewPublicDnsNamespace(stack, jsii.String("Namespace"), &PublicDnsNamespaceProps{
//   	Name: jsii.String("foobar.com"),
//   })
//
//   service := namespace.CreateService(jsii.String("Service"), &DnsServiceProps{
//   	Name: jsii.String("foo"),
//   	DnsRecordType: servicediscovery.DnsRecordType_A,
//   	DnsTtl: cdk.Duration_Seconds(jsii.Number(30)),
//   	HealthCheck: &HealthCheckConfig{
//   		Type: servicediscovery.HealthCheckType_HTTPS,
//   		ResourcePath: jsii.String("/healthcheck"),
//   		FailureThreshold: jsii.Number(2),
//   	},
//   })
//
//   service.RegisterIpInstance(jsii.String("IpInstance"), &IpInstanceBaseProps{
//   	Ipv4: jsii.String("54.239.25.192"),
//   	Port: jsii.Number(443),
//   })
//
//   app.Synth()
//
type PublicDnsNamespaceProps struct {
	// A name for the Namespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the Namespace.
	// Default: none.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

