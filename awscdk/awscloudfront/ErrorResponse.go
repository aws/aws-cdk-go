package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for configuring custom error responses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   errorResponse := &ErrorResponse{
//   	HttpStatus: jsii.Number(123),
//
//   	// the properties below are optional
//   	ResponseHttpStatus: jsii.Number(123),
//   	ResponsePagePath: jsii.String("responsePagePath"),
//   	Ttl: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type ErrorResponse struct {
	// The HTTP status code for which you want to specify a custom error page and/or a caching duration.
	HttpStatus *float64 `field:"required" json:"httpStatus" yaml:"httpStatus"`
	// The HTTP status code that you want CloudFront to return to the viewer along with the custom error page.
	//
	// If you specify a value for `responseHttpStatus`, you must also specify a value for `responsePagePath`.
	// Default: - the error code will be returned as the response code.
	//
	ResponseHttpStatus *float64 `field:"optional" json:"responseHttpStatus" yaml:"responseHttpStatus"`
	// The path to the custom error page that you want CloudFront to return to a viewer when your origin returns the `httpStatus`, for example, /4xx-errors/403-forbidden.html.
	// Default: - the default CloudFront response is shown.
	//
	ResponsePagePath *string `field:"optional" json:"responsePagePath" yaml:"responsePagePath"`
	// The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status code specified in ErrorCode.
	// Default: - the default caching TTL behavior applies.
	//
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
}

