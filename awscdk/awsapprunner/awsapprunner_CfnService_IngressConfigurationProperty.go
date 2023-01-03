package awsapprunner


// Network configuration settings for inbound network traffic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressConfigurationProperty := &ingressConfigurationProperty{
//   	isPubliclyAccessible: jsii.Boolean(false),
//   }
//
type CfnService_IngressConfigurationProperty struct {
	// Specifies whether your App Runner service is publicly accessible.
	//
	// To make the service publicly accessible set it to `True` . To make the service privately accessible, from only within an Amazon VPC set it to `False` .
	IsPubliclyAccessible interface{} `field:"required" json:"isPubliclyAccessible" yaml:"isPubliclyAccessible"`
}

