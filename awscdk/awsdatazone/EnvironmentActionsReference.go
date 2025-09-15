package awsdatazone


// A reference to a EnvironmentActions resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentActionsReference := &EnvironmentActionsReference{
//   	DomainId: jsii.String("domainId"),
//   	EnvironmentActionsId: jsii.String("environmentActionsId"),
//   	EnvironmentId: jsii.String("environmentId"),
//   }
//
type EnvironmentActionsReference struct {
	// The DomainId of the EnvironmentActions resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The Id of the EnvironmentActions resource.
	EnvironmentActionsId *string `field:"required" json:"environmentActionsId" yaml:"environmentActionsId"`
	// The EnvironmentId of the EnvironmentActions resource.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
}

