package awsvpclattice


// The codes to use when checking for a successful response from a target for health checks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matcherProperty := &MatcherProperty{
//   	HttpCode: jsii.String("httpCode"),
//   }
//
type CfnTargetGroup_MatcherProperty struct {
	// The HTTP code to use when checking for a successful response from a target.
	HttpCode *string `field:"required" json:"httpCode" yaml:"httpCode"`
}

