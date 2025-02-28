package awscloudfront


// A complex type that controls:.
//
// - Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.
// - How long CloudFront caches HTTP status codes in the 4xx and 5xx range.
//
// For more information about custom error pages, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customErrorResponseProperty := &CustomErrorResponseProperty{
//   	ErrorCode: jsii.Number(123),
//
//   	// the properties below are optional
//   	ErrorCachingMinTtl: jsii.Number(123),
//   	ResponseCode: jsii.Number(123),
//   	ResponsePagePath: jsii.String("responsePagePath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customerrorresponse.html
//
type CfnDistribution_CustomErrorResponseProperty struct {
	// The HTTP status code for which you want to specify a custom error page and/or a caching duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customerrorresponse.html#cfn-cloudfront-distribution-customerrorresponse-errorcode
	//
	ErrorCode *float64 `field:"required" json:"errorCode" yaml:"errorCode"`
	// The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status code specified in `ErrorCode` .
	//
	// When this time period has elapsed, CloudFront queries your origin to see whether the problem that caused the error has been resolved and the requested object is now available.
	//
	// For more information, see [Customizing Error Responses](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customerrorresponse.html#cfn-cloudfront-distribution-customerrorresponse-errorcachingminttl
	//
	// Default: - 300.
	//
	ErrorCachingMinTtl *float64 `field:"optional" json:"errorCachingMinTtl" yaml:"errorCachingMinTtl"`
	// The HTTP status code that you want CloudFront to return to the viewer along with the custom error page.
	//
	// There are a variety of reasons that you might want CloudFront to return a status code different from the status code that your origin returned to CloudFront, for example:
	//
	// - Some Internet devices (some firewalls and corporate proxies, for example) intercept HTTP 4xx and 5xx and prevent the response from being returned to the viewer. If you substitute `200` , the response typically won't be intercepted.
	// - If you don't care about distinguishing among different client errors or server errors, you can specify `400` or `500` as the `ResponseCode` for all 4xx or 5xx errors.
	// - You might want to return a `200` status code (OK) and static website so your customers don't know that your website is down.
	//
	// If you specify a value for `ResponseCode` , you must also specify a value for `ResponsePagePath` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customerrorresponse.html#cfn-cloudfront-distribution-customerrorresponse-responsecode
	//
	ResponseCode *float64 `field:"optional" json:"responseCode" yaml:"responseCode"`
	// The path to the custom error page that you want CloudFront to return to a viewer when your origin returns the HTTP status code specified by `ErrorCode` , for example, `/4xx-errors/403-forbidden.html` . If you want to store your objects and your custom error pages in different locations, your distribution must include a cache behavior for which the following is true:.
	//
	// - The value of `PathPattern` matches the path to your custom error messages. For example, suppose you saved custom error pages for 4xx errors in an Amazon S3 bucket in a directory named `/4xx-errors` . Your distribution must include a cache behavior for which the path pattern routes requests for your custom error pages to that location, for example, `/4xx-errors/*` .
	// - The value of `TargetOriginId` specifies the value of the `ID` element for the origin that contains your custom error pages.
	//
	// If you specify a value for `ResponsePagePath` , you must also specify a value for `ResponseCode` .
	//
	// We recommend that you store custom error pages in an Amazon S3 bucket. If you store custom error pages on an HTTP server and the server starts to return 5xx errors, CloudFront can't get the files that you want to return to viewers because the origin server is unavailable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-customerrorresponse.html#cfn-cloudfront-distribution-customerrorresponse-responsepagepath
	//
	ResponsePagePath *string `field:"optional" json:"responsePagePath" yaml:"responsePagePath"`
}

