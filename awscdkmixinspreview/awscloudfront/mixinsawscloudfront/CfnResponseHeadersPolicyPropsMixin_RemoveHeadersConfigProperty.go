package mixinsawscloudfront


// A list of HTTP header names that CloudFront removes from HTTP responses to requests that match the cache behavior that this response headers policy is attached to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   removeHeadersConfigProperty := &RemoveHeadersConfigProperty{
//   	Items: []interface{}{
//   		&RemoveHeaderProperty{
//   			Header: jsii.String("header"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-removeheadersconfig.html
//
type CfnResponseHeadersPolicyPropsMixin_RemoveHeadersConfigProperty struct {
	// The list of HTTP header names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-responseheaderspolicy-removeheadersconfig.html#cfn-cloudfront-responseheaderspolicy-removeheadersconfig-items
	//
	Items interface{} `field:"optional" json:"items" yaml:"items"`
}

