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
//   var v1 interface{}
//   var v2 interface{}
//
//   featureFlag := &FeatureFlag{
//   	Explanation: jsii.String("explanation"),
//   	RecommendedValue: recommendedValue,
//   	UnconfiguredBehavesLike: &UnconfiguredBehavesLike{
//   		V1: v1,
//   		V2: v2,
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
	// The value of the flag that produces the same behavior as when the flag is not configured at all.
	//
	// The structure of this field is a historical accident. The type of this field
	// should have been boolean, which should have contained the default value for
	// the flag appropriate for the *current* version of the CDK library. We are
	// not rectifying this accident because doing so
	//
	// Instead, the canonical way to access this value is by evaluating
	// `unconfiguredBehavesLike?.v2 ?? false`.
	// Default: false.
	//
	UnconfiguredBehavesLike *UnconfiguredBehavesLike `field:"optional" json:"unconfiguredBehavesLike" yaml:"unconfiguredBehavesLike"`
	// The value configured by the user.
	//
	// This is the value configured at the root of the tree. Users may also have
	// configured values at specific locations in the tree; we don't report on
	// those.
	// Default: - Not configured by the user.
	//
	UserValue interface{} `field:"optional" json:"userValue" yaml:"userValue"`
}

