package awsappmesh


// Configuration for `HeaderMatch`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerMatchConfig := &headerMatchConfig{
//   	headerMatch: &httpRouteHeaderProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		invert: jsii.Boolean(false),
//   		match: &headerMatchMethodProperty{
//   			exact: jsii.String("exact"),
//   			prefix: jsii.String("prefix"),
//   			range: &matchRangeProperty{
//   				end: jsii.Number(123),
//   				start: jsii.Number(123),
//   			},
//   			regex: jsii.String("regex"),
//   			suffix: jsii.String("suffix"),
//   		},
//   	},
//   }
//
type HeaderMatchConfig struct {
	// Route CFN configuration for the route header match.
	HeaderMatch *CfnRoute_HttpRouteHeaderProperty `field:"required" json:"headerMatch" yaml:"headerMatch"`
}

