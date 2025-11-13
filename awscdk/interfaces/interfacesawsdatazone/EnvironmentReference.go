package interfacesawsdatazone


// A reference to a Environment resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentReference := &EnvironmentReference{
//   	DomainId: jsii.String("domainId"),
//   	EnvironmentId: jsii.String("environmentId"),
//   }
//
type EnvironmentReference struct {
	// The DomainId of the Environment resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The Id of the Environment resource.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
}

