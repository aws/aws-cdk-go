package awsappmesh


// The type returned from the `bind()` method in {@link HttpRoutePathMatch}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRoutePathMatchConfig := &httpRoutePathMatchConfig{
//   	prefixPathMatch: jsii.String("prefixPathMatch"),
//   	wholePathMatch: &httpPathMatchProperty{
//   		exact: jsii.String("exact"),
//   		regex: jsii.String("regex"),
//   	},
//   }
//
type HttpRoutePathMatchConfig struct {
	// Route configuration for matching on the prefix of the URL path of the request.
	PrefixPathMatch *string `field:"optional" json:"prefixPathMatch" yaml:"prefixPathMatch"`
	// Route configuration for matching on the complete URL path of the request.
	WholePathMatch *CfnRoute_HttpPathMatchProperty `field:"optional" json:"wholePathMatch" yaml:"wholePathMatch"`
}

