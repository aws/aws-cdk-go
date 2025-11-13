package interfacesawscontroltower


// A reference to a EnabledBaseline resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enabledBaselineReference := &EnabledBaselineReference{
//   	EnabledBaselineIdentifier: jsii.String("enabledBaselineIdentifier"),
//   }
//
type EnabledBaselineReference struct {
	// The EnabledBaselineIdentifier of the EnabledBaseline resource.
	EnabledBaselineIdentifier *string `field:"required" json:"enabledBaselineIdentifier" yaml:"enabledBaselineIdentifier"`
}

