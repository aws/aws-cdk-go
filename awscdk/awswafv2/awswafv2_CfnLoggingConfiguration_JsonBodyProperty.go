package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   jsonBodyProperty := &jsonBodyProperty{
//   	matchPattern: &matchPatternProperty{
//   		all: all,
//   		includedPaths: []*string{
//   			jsii.String("includedPaths"),
//   		},
//   	},
//   	matchScope: jsii.String("matchScope"),
//
//   	// the properties below are optional
//   	invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   }
//
type CfnLoggingConfiguration_JsonBodyProperty struct {
	// `CfnLoggingConfiguration.JsonBodyProperty.MatchPattern`.
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// `CfnLoggingConfiguration.JsonBodyProperty.MatchScope`.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// `CfnLoggingConfiguration.JsonBodyProperty.InvalidFallbackBehavior`.
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
}

