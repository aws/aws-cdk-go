package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   jsonBodyProperty := &JsonBodyProperty{
//   	MatchPattern: &MatchPatternProperty{
//   		All: all,
//   		IncludedPaths: []*string{
//   			jsii.String("includedPaths"),
//   		},
//   	},
//   	MatchScope: jsii.String("matchScope"),
//
//   	// the properties below are optional
//   	InvalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-jsonbody.html
//
type CfnLoggingConfiguration_JsonBodyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-jsonbody.html#cfn-wafv2-loggingconfiguration-jsonbody-matchpattern
	//
	MatchPattern interface{} `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-jsonbody.html#cfn-wafv2-loggingconfiguration-jsonbody-matchscope
	//
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-jsonbody.html#cfn-wafv2-loggingconfiguration-jsonbody-invalidfallbackbehavior
	//
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
}

