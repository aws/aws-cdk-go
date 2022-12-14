package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//
//   matchPatternProperty := &matchPatternProperty{
//   	all: all,
//   	includedPaths: []*string{
//   		jsii.String("includedPaths"),
//   	},
//   }
//
type CfnLoggingConfiguration_MatchPatternProperty struct {
	// `CfnLoggingConfiguration.MatchPatternProperty.All`.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// `CfnLoggingConfiguration.MatchPatternProperty.IncludedPaths`.
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

