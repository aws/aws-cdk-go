package mixinsawscloudfront


// A list of HTTP response header names and their values.
//
// CloudFront includes these headers in HTTP responses that it sends for requests that match a cache behavior that's associated with this response headers policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customHeadersConfigProperty := &CustomHeadersConfigProperty{
//   	Items: []interface{}{
//   		&CustomHeaderProperty{
//   			Header: jsii.String("header"),
//   			Override: jsii.Boolean(false),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-customheadersconfig.html
//
type CfnResponseHeadersPolicyPropsMixin_CustomHeadersConfigProperty struct {
	// The list of HTTP response headers and their values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-customheadersconfig.html#cfn-cloudfront-responseheaderspolicy-customheadersconfig-items
	//
	Items interface{} `field:"optional" json:"items" yaml:"items"`
}

