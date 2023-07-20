package awselasticloadbalancingv2


// Specifies the HTTP codes that healthy targets must use when responding to an HTTP health check.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matcherProperty := &MatcherProperty{
//   	GrpcCode: jsii.String("grpcCode"),
//   	HttpCode: jsii.String("httpCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html
//
type CfnTargetGroup_MatcherProperty struct {
	// You can specify values between 0 and 99.
	//
	// You can specify multiple values (for example, "0,1") or a range of values (for example, "0-5"). The default value is 12.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html#cfn-elasticloadbalancingv2-targetgroup-matcher-grpccode
	//
	GrpcCode *string `field:"optional" json:"grpcCode" yaml:"grpcCode"`
	// For Application Load Balancers, you can specify values between 200 and 499, with the default value being 200.
	//
	// You can specify multiple values (for example, "200,202") or a range of values (for example, "200-299").
	//
	// For Network Load Balancers, you can specify values between 200 and 599, with the default value being 200-399. You can specify multiple values (for example, "200,202") or a range of values (for example, "200-299").
	//
	// For Gateway Load Balancers, this must be "200–399".
	//
	// Note that when using shorthand syntax, some values such as commas need to be escaped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html#cfn-elasticloadbalancingv2-targetgroup-matcher-httpcode
	//
	HttpCode *string `field:"optional" json:"httpCode" yaml:"httpCode"`
}

