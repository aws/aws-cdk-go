package awslightsail


// `CookieObject` is a property of the [CacheSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-cachesettings.html) property. It describes whether an Amazon Lightsail content delivery network (CDN) distribution forwards cookies to the origin and, if so, which ones.
//
// For the cookies that you specify, your distribution caches separate versions of the specified content based on the cookie values in viewer requests.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cookieObjectProperty := &cookieObjectProperty{
//   	cookiesAllowList: []*string{
//   		jsii.String("cookiesAllowList"),
//   	},
//   	option: jsii.String("option"),
//   }
//
type CfnDistribution_CookieObjectProperty struct {
	// The specific cookies to forward to your distribution's origin.
	CookiesAllowList *[]*string `field:"optional" json:"cookiesAllowList" yaml:"cookiesAllowList"`
	// Specifies which cookies to forward to the distribution's origin for a cache behavior.
	//
	// Use one of the following configurations for your distribution:
	//
	// - *`all`* - Forwards all cookies to your origin.
	// - *`none`* - Doesn’t forward cookies to your origin.
	// - *`allow-list`* - Forwards only the cookies that you specify using the `CookiesAllowList` parameter.
	Option *string `field:"optional" json:"option" yaml:"option"`
}

