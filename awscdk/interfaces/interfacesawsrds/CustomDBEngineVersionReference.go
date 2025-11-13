package interfacesawsrds


// A reference to a CustomDBEngineVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customDBEngineVersionReference := &CustomDBEngineVersionReference{
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   }
//
type CustomDBEngineVersionReference struct {
	// The Engine of the CustomDBEngineVersion resource.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The EngineVersion of the CustomDBEngineVersion resource.
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
}

