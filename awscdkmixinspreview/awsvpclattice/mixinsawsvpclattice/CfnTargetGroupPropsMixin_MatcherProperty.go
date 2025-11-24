package mixinsawsvpclattice


// Describes the codes to use when checking for a successful response from a target for health checks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   matcherProperty := &MatcherProperty{
//   	HttpCode: jsii.String("httpCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-matcher.html
//
type CfnTargetGroupPropsMixin_MatcherProperty struct {
	// The HTTP code to use when checking for a successful response from a target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-targetgroup-matcher.html#cfn-vpclattice-targetgroup-matcher-httpcode
	//
	HttpCode *string `field:"optional" json:"httpCode" yaml:"httpCode"`
}

