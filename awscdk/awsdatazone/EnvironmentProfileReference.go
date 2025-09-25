package awsdatazone


// A reference to a EnvironmentProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentProfileReference := &EnvironmentProfileReference{
//   	DomainId: jsii.String("domainId"),
//   	EnvironmentProfileId: jsii.String("environmentProfileId"),
//   }
//
type EnvironmentProfileReference struct {
	// The DomainId of the EnvironmentProfile resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The Id of the EnvironmentProfile resource.
	EnvironmentProfileId *string `field:"required" json:"environmentProfileId" yaml:"environmentProfileId"`
}

