package awsappmesh


// The type returned from the `bind()` method in `HttpRoutePathMatch`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   httpRoutePathMatchConfig := &HttpRoutePathMatchConfig{
//   	PrefixPathMatch: jsii.String("prefixPathMatch"),
//   	WholePathMatch: &HttpPathMatchProperty{
//   		Exact: jsii.String("exact"),
//   		Regex: jsii.String("regex"),
//   	},
//   }
//
type HttpRoutePathMatchConfig struct {
	// Route configuration for matching on the prefix of the URL path of the request.
	// Default: - no matching will be performed on the prefix of the URL path.
	//
	PrefixPathMatch *string `field:"optional" json:"prefixPathMatch" yaml:"prefixPathMatch"`
	// Route configuration for matching on the complete URL path of the request.
	// Default: - no matching will be performed on the complete URL path.
	//
	WholePathMatch *CfnRoute_HttpPathMatchProperty `field:"optional" json:"wholePathMatch" yaml:"wholePathMatch"`
}

