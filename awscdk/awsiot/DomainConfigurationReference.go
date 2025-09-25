package awsiot


// A reference to a DomainConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainConfigurationReference := &DomainConfigurationReference{
//   	DomainConfigurationArn: jsii.String("domainConfigurationArn"),
//   	DomainConfigurationName: jsii.String("domainConfigurationName"),
//   }
//
type DomainConfigurationReference struct {
	// The ARN of the DomainConfiguration resource.
	DomainConfigurationArn *string `field:"required" json:"domainConfigurationArn" yaml:"domainConfigurationArn"`
	// The DomainConfigurationName of the DomainConfiguration resource.
	DomainConfigurationName *string `field:"required" json:"domainConfigurationName" yaml:"domainConfigurationName"`
}

