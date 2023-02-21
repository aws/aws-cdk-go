package awslightsail


// `HeaderObject` is a property of the [CacheSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-cachesettings.html) property. It describes the request headers used by your distribution, which caches your content based on the request headers.
//
// For the headers that you specify, your distribution caches separate versions of the specified content based on the header values in viewer requests. For example, suppose that viewer requests for logo.jpg contain a custom product header that has a value of either acme or apex. Also, suppose that you configure your distribution to cache your content based on values in the product header. Your distribution forwards the product header to the origin and caches the response from the origin once for each header value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   headerObjectProperty := &headerObjectProperty{
//   	headersAllowList: []*string{
//   		jsii.String("headersAllowList"),
//   	},
//   	option: jsii.String("option"),
//   }
//
type CfnDistribution_HeaderObjectProperty struct {
	// The specific headers to forward to your distribution's origin.
	HeadersAllowList *[]*string `field:"optional" json:"headersAllowList" yaml:"headersAllowList"`
	// The headers that you want your distribution to forward to your origin.
	//
	// Your distribution caches your content based on these headers.
	//
	// Use one of the following configurations for your distribution:
	//
	// - *`all`* - Forwards all headers to your origin..
	// - *`none`* - Forwards only the default headers.
	// - *`allow-list`* - Forwards only the headers that you specify using the `HeadersAllowList` parameter.
	Option *string `field:"optional" json:"option" yaml:"option"`
}

