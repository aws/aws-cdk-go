package awscdkapplicationsignalsalpha


// Interface for environment extensions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import applicationsignals_alpha "github.com/aws/aws-cdk-go/awscdkapplicationsignalsalpha"
//
//   environmentExtension := &EnvironmentExtension{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// Experimental.
type EnvironmentExtension struct {
	// The name of the environment variable.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the environment variable.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

