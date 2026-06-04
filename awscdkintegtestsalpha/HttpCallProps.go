package awscdkintegtestsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Options for creating an HttpApiCall provider.
//
// Example:
//   var stack Stack
//
//
//   awscdkintegtestsalpha.NewHttpApiCall(stack, jsii.String("MyAsssertion"), &HttpCallProps{
//   	Url: jsii.String("https://example-api.com/abc"),
//   })
//
// Experimental.
type HttpCallProps struct {
	// The url to fetch.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Options for fetch.
	// Experimental.
	FetchOptions *FetchOptions `field:"optional" json:"fetchOptions" yaml:"fetchOptions"`
	// The log level of the provider lambda function.
	// Default: ApplicationLogLevel.FATAL
	//
	// Experimental.
	ProviderLogLevel awslambda.ApplicationLogLevel `field:"optional" json:"providerLogLevel" yaml:"providerLogLevel"`
}

