package awscdkintegtestsalpha


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   httpRequestParameters := &HttpRequestParameters{
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	FetchOptions: &FetchOptions{
//   		Body: jsii.String("body"),
//   		Headers: map[string]*string{
//   			"headersKey": jsii.String("headers"),
//   		},
//   		Method: jsii.String("method"),
//   		Port: jsii.Number(123),
//   	},
//   }
//
// Experimental.
type HttpRequestParameters struct {
	// The url to fetch.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Options for fetch.
	// Experimental.
	FetchOptions *FetchOptions `field:"optional" json:"fetchOptions" yaml:"fetchOptions"`
}

