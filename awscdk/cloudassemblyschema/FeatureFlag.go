package cloudassemblyschema


// A single feature flag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var recommendedValue interface{}
//   var unconfiguredBehavesLike interface{}
//   var userValue interface{}
//
//   featureFlag := &FeatureFlag{
//   	Explanation: jsii.String("explanation"),
//   	RecommendedValue: recommendedValue,
//   	UnconfiguredBehavesLike: map[string]interface{}{
//   		"unconfiguredBehavesLikeKey": unconfiguredBehavesLike,
//   	},
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
	// The value of the flag if it is unconfigured.
	// Default: - No value.
	//
	UnconfiguredBehavesLike *map[string]interface{} `field:"optional" json:"unconfiguredBehavesLike" yaml:"unconfiguredBehavesLike"`
	// The value configured by the user.
	//
	// This is the value configured at the root of the tree. Users may also have
	// configured values at specific locations in the tree; we don't report on
	// those.
	// Default: - Not configured by the user.
	//
	UserValue interface{} `field:"optional" json:"userValue" yaml:"userValue"`
}

