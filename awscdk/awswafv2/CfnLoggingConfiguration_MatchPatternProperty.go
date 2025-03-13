package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   matchPatternProperty := &MatchPatternProperty{
//   	All: all,
//   	IncludedPaths: []*string{
//   		jsii.String("includedPaths"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-matchpattern.html
//
type CfnLoggingConfiguration_MatchPatternProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-matchpattern.html#cfn-wafv2-loggingconfiguration-matchpattern-all
	//
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-loggingconfiguration-matchpattern.html#cfn-wafv2-loggingconfiguration-matchpattern-includedpaths
	//
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

