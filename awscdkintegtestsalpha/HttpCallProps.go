package awscdkintegtestsalpha


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
}

