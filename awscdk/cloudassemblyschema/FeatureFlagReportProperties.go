package cloudassemblyschema


// Artifact properties for a feature flag report.
//
// A feature flag report is small enough that all the properties can be inlined
// here, and doesn't need an additional file.
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
//   featureFlagReportProperties := &FeatureFlagReportProperties{
//   	Flags: map[string]FeatureFlag{
//   		"flagsKey": &FeatureFlag{
//   			"explanation": jsii.String("explanation"),
//   			"recommendedValue": recommendedValue,
//   			"unconfiguredBehavesLike": map[string]interface{}{
//   				"unconfiguredBehavesLikeKey": unconfiguredBehavesLike,
//   			},
//   			"userValue": userValue,
//   		},
//   	},
//   	Module: jsii.String("module"),
//   }
//
type FeatureFlagReportProperties struct {
	// Information about every feature flag supported by this library.
	Flags *map[string]*FeatureFlag `field:"required" json:"flags" yaml:"flags"`
	// The library that this feature flag report applies to.
	Module *string `field:"required" json:"module" yaml:"module"`
}

