package interfacesawsdatazone


// A reference to a EnvironmentBlueprintConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentBlueprintConfigurationReference := &EnvironmentBlueprintConfigurationReference{
//   	DomainId: jsii.String("domainId"),
//   	EnvironmentBlueprintId: jsii.String("environmentBlueprintId"),
//   }
//
type EnvironmentBlueprintConfigurationReference struct {
	// The DomainId of the EnvironmentBlueprintConfiguration resource.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The EnvironmentBlueprintId of the EnvironmentBlueprintConfiguration resource.
	EnvironmentBlueprintId *string `field:"required" json:"environmentBlueprintId" yaml:"environmentBlueprintId"`
}

