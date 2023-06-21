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
type HealthCheckType string

const (
	// Route 53 tries to establish a TCP connection.
	//
	// If successful, Route 53 submits an HTTP request and waits for an HTTP
	// status code of 200 or greater and less than 400.
	HealthCheckType_HTTP HealthCheckType = "HTTP"
	// Route 53 tries to establish a TCP connection.
	//
	// If successful, Route 53 submits an HTTPS request and waits for an
	// HTTP status code of 200 or greater and less than 400.  If you specify HTTPS for the value of Type, the endpoint
	// must support TLS v1.0 or later.
	HealthCheckType_HTTPS HealthCheckType = "HTTPS"
	// Route 53 tries to establish a TCP connection.
	//
	// If you specify TCP for Type, don't specify a value for ResourcePath.
	HealthCheckType_TCP HealthCheckType = "TCP"
)

