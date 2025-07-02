package cloudassemblyschema


// A single feature flag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var recommendedValue interface{}
//   var userValue interface{}
//
//   featureFlag := &FeatureFlag{
//   	Explanation: jsii.String("explanation"),
//   	RecommendedValue: recommendedValue,
//   	UserValue: userValue,
//   }
//
type FeatureFlag struct {
	// Explanation about the purpose of this flag that can be shown to the user.
	// Default: - No description.
	//
	Explanation *string `field:"optional" json:"explanation" yaml:"explanation"`
	// The library-recommended value for this flag, if any.
	//
	// It is possible that there is no recommended value.
	// Default: - No recommended value.
	//
	RecommendedValue interface{} `field:"optional" json:"recommendedValue" yaml:"recommendedValue"`
	// The value configured by the user.
	//
	// This is the value configured at the root of the tree. Users may also have
	// configured values at specific locations in the tree; we don't report on
	// those.
	// Default: - Not configured by the user.
	//
	UserValue interface{} `field:"optional" json:"userValue" yaml:"userValue"`
}

