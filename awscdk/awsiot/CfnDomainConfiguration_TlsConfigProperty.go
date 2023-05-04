package awsiot


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tlsConfigProperty := &TlsConfigProperty{
//   	SecurityPolicy: jsii.String("securityPolicy"),
//   }
//
type CfnDomainConfiguration_TlsConfigProperty struct {
	// `CfnDomainConfiguration.TlsConfigProperty.SecurityPolicy`.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}

