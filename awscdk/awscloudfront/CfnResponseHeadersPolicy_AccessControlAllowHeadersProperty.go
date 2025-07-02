package awscloudfront


// A list of HTTP header names that CloudFront includes as values for the `Access-Control-Allow-Headers` HTTP response header.
//
// For more information about the `Access-Control-Allow-Headers` HTTP response header, see [Access-Control-Allow-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlAllowHeadersProperty := &AccessControlAllowHeadersProperty{
//   	Items: []*string{
//   		jsii.String("items"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-accesscontrolallowheaders.html
//
type CfnResponseHeadersPolicy_AccessControlAllowHeadersProperty struct {
	// The list of HTTP header names.
	//
	// You can specify `*` to allow all headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-accesscontrolallowheaders.html#cfn-cloudfront-responseheaderspolicy-accesscontrolallowheaders-items
	//
	Items *[]*string `field:"required" json:"items" yaml:"items"`
}

