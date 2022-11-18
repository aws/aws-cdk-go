package awsapprunner


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
	// `CfnService.IngressConfigurationProperty.IsPubliclyAccessible`.
	IsPubliclyAccessible interface{} `field:"required" json:"isPubliclyAccessible" yaml:"isPubliclyAccessible"`
}

