package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainSAMPTProperty := &domainSAMPTProperty{
//   	domainName: jsii.String("domainName"),
//   }
//
type CfnFunction_DomainSAMPTProperty struct {
	// `CfnFunction.DomainSAMPTProperty.DomainName`.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

