package awsappmesh


// Configuration for `HeaderMatch`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerMatchConfig := &HeaderMatchConfig{
//   	HeaderMatch: &HttpRouteHeaderProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Invert: jsii.Boolean(false),
//   		Match: &HeaderMatchMethodProperty{
//   			Exact: jsii.String("exact"),
//   			Prefix: jsii.String("prefix"),
//   			Range: &MatchRangeProperty{
//   				End: jsii.Number(123),
//   				Start: jsii.Number(123),
//   			},
//   			Regex: jsii.String("regex"),
//   			Suffix: jsii.String("suffix"),
//   		},
//   	},
//   }
//
// Experimental.
type HeaderMatchConfig struct {
	// Route CFN configuration for the route header match.
	// Experimental.
	HeaderMatch *CfnRoute_HttpRouteHeaderProperty `field:"required" json:"headerMatch" yaml:"headerMatch"`
}

