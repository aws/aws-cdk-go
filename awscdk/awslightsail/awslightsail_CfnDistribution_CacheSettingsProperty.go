package awslightsail


// `CacheSettings` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the cache settings of an Amazon Lightsail content delivery network (CDN) distribution.
//
// These settings apply only to your distribution’s `CacheBehaviors` that have a `Behavior` of `cache` . This includes the `DefaultCacheBehavior` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheSettingsProperty := &cacheSettingsProperty{
//   	allowedHttpMethods: jsii.String("allowedHttpMethods"),
//   	cachedHttpMethods: jsii.String("cachedHttpMethods"),
//   	defaultTtl: jsii.Number(123),
//   	forwardedCookies: &cookieObjectProperty{
//   		cookiesAllowList: []*string{
//   			jsii.String("cookiesAllowList"),
//   		},
//   		option: jsii.String("option"),
//   	},
//   	forwardedHeaders: &headerObjectProperty{
//   		headersAllowList: []*string{
//   			jsii.String("headersAllowList"),
//   		},
//   		option: jsii.String("option"),
//   	},
//   	forwardedQueryStrings: &queryStringObjectProperty{
//   		option: jsii.Boolean(false),
//   		queryStringsAllowList: []*string{
//   			jsii.String("queryStringsAllowList"),
//   		},
//   	},
//   	maximumTtl: jsii.Number(123),
//   	minimumTtl: jsii.Number(123),
//   }
//
type CfnDistribution_CacheSettingsProperty struct {
	// The HTTP methods that are processed and forwarded to the distribution's origin.
	//
	// You can specify the following options:
	//
	// - `GET,HEAD` - The distribution forwards the `GET` and `HEAD` methods.
	// - `GET,HEAD,OPTIONS` - The distribution forwards the `GET` , `HEAD` , and `OPTIONS` methods.
	// - `GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE` - The distribution forwards the `GET` , `HEAD` , `OPTIONS` , `PUT` , `PATCH` , `POST` , and `DELETE` methods.
	//
	// If you specify `GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE` , you might need to restrict access to your distribution's origin so users can't perform operations that you don't want them to. For example, you might not want users to have permission to delete objects from your origin.
	AllowedHttpMethods *string `field:"optional" json:"allowedHttpMethods" yaml:"allowedHttpMethods"`
	// The HTTP method responses that are cached by your distribution.
	//
	// You can specify the following options:
	//
	// - `GET,HEAD` - The distribution caches responses to the `GET` and `HEAD` methods.
	// - `GET,HEAD,OPTIONS` - The distribution caches responses to the `GET` , `HEAD` , and `OPTIONS` methods.
	CachedHttpMethods *string `field:"optional" json:"cachedHttpMethods" yaml:"cachedHttpMethods"`
	// The default amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the content has been updated.
	//
	// > The value specified applies only when the origin does not add HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects.
	DefaultTtl *float64 `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// An object that describes the cookies that are forwarded to the origin.
	//
	// Your content is cached based on the cookies that are forwarded.
	ForwardedCookies interface{} `field:"optional" json:"forwardedCookies" yaml:"forwardedCookies"`
	// An object that describes the headers that are forwarded to the origin.
	//
	// Your content is cached based on the headers that are forwarded.
	ForwardedHeaders interface{} `field:"optional" json:"forwardedHeaders" yaml:"forwardedHeaders"`
	// An object that describes the query strings that are forwarded to the origin.
	//
	// Your content is cached based on the query strings that are forwarded.
	ForwardedQueryStrings interface{} `field:"optional" json:"forwardedQueryStrings" yaml:"forwardedQueryStrings"`
	// The maximum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// The value specified applies only when the origin adds HTTP headers such as `Cache-Control max-age` , `Cache-Control s-maxage` , and `Expires` to objects.
	MaximumTtl *float64 `field:"optional" json:"maximumTtl" yaml:"maximumTtl"`
	// The minimum amount of time that objects stay in the distribution's cache before the distribution forwards another request to the origin to determine whether the object has been updated.
	//
	// A value of `0` must be specified for `minimumTTL` if the distribution is configured to forward all headers to the origin.
	MinimumTtl *float64 `field:"optional" json:"minimumTtl" yaml:"minimumTtl"`
}

