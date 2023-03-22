package awsvpclattice


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
	// `CfnTargetGroup.MatcherProperty.HttpCode`.
	HttpCode *string `field:"required" json:"httpCode" yaml:"httpCode"`
}

