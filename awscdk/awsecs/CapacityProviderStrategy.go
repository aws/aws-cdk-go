package awsecs


// A Capacity Provider strategy to use for the service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityProviderStrategy := &CapacityProviderStrategy{
//   	CapacityProvider: jsii.String("capacityProvider"),
//
//   	// the properties below are optional
//   	Base: jsii.Number(123),
//   	Weight: jsii.Number(123),
//   }
//
type CapacityProviderStrategy struct {
	// The name of the capacity provider.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// The base value designates how many tasks, at a minimum, to run on the specified capacity provider.
	//
	// Only one
	// capacity provider in a capacity provider strategy can have a base defined. If no value is specified, the default
	// value of 0 is used.
	// Default: - none.
	//
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// The weight value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider.
	//
	// The weight value is taken into consideration after the base value, if defined, is satisfied.
	// Default: - 0.
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

