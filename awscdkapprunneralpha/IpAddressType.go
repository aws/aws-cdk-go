package awscdkapprunneralpha


// The IP address type for your incoming public network configuration.
//
// Example:
//   apprunner.NewService(this, jsii.String("Service"), &ServiceProps{
//   	Source: apprunner.Source_FromEcrPublic(&EcrPublicProps{
//   		ImageConfiguration: &ImageConfiguration{
//   			Port: jsii.Number(8000),
//   		},
//   		ImageIdentifier: jsii.String("public.ecr.aws/aws-containers/hello-app-runner:latest"),
//   	}),
//   	IpAddressType: apprunner.IpAddressType_DUAL_STACK,
//   })
//
// Experimental.
type IpAddressType string

const (
	// IPV4.
	// Experimental.
	IpAddressType_IPV4 IpAddressType = "IPV4"
	// DUAL_STACK.
	// Experimental.
	IpAddressType_DUAL_STACK IpAddressType = "DUAL_STACK"
)

