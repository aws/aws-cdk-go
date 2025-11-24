package mixinsawscloudfront


// An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in requests that CloudFront sends to the origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   queryStringsConfigProperty := &QueryStringsConfigProperty{
//   	QueryStringBehavior: jsii.String("queryStringBehavior"),
//   	QueryStrings: []*string{
//   		jsii.String("queryStrings"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-querystringsconfig.html
//
type CfnOriginRequestPolicyPropsMixin_QueryStringsConfigProperty struct {
	// Determines whether any URL query strings in viewer requests are included in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//
	// - `none` – No query strings in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any query strings that are listed in a `CachePolicy` *are* included in origin requests.
	// - `whitelist` – Only the query strings in viewer requests that are listed in the `QueryStringNames` type are included in requests that CloudFront sends to the origin.
	// - `all` – All query strings in viewer requests are included in requests that CloudFront sends to the origin.
	// - `allExcept` – All query strings in viewer requests are included in requests that CloudFront sends to the origin, **except** for those listed in the `QueryStringNames` type, which are not included.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-querystringsconfig.html#cfn-cloudfront-originrequestpolicy-querystringsconfig-querystringbehavior
	//
	QueryStringBehavior *string `field:"optional" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// Contains a list of query string names.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-originrequestpolicy-querystringsconfig.html#cfn-cloudfront-originrequestpolicy-querystringsconfig-querystrings
	//
	QueryStrings *[]*string `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

