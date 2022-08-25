package awscloudfront


// A configuration for a set of HTTP response headers that are used for cross-origin resource sharing (CORS).
//
// CloudFront adds these headers to HTTP responses that it sends for CORS requests that match a cache behavior associated with this response headers policy.
//
// For more information about CORS, see [Cross-Origin Resource Sharing (CORS)](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsConfigProperty := &corsConfigProperty{
//   	accessControlAllowCredentials: jsii.Boolean(false),
//   	accessControlAllowHeaders: &accessControlAllowHeadersProperty{
//   		items: []*string{
//   			jsii.String("items"),
//   		},
//   	},
//   	accessControlAllowMethods: &accessControlAllowMethodsProperty{
//   		items: []*string{
//   			jsii.String("items"),
//   		},
//   	},
//   	accessControlAllowOrigins: &accessControlAllowOriginsProperty{
//   		items: []*string{
//   			jsii.String("items"),
//   		},
//   	},
//   	originOverride: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	accessControlExposeHeaders: &accessControlExposeHeadersProperty{
//   		items: []*string{
//   			jsii.String("items"),
//   		},
//   	},
//   	accessControlMaxAgeSec: jsii.Number(123),
//   }
//
type CfnResponseHeadersPolicy_CorsConfigProperty struct {
	// A Boolean that CloudFront uses as the value for the `Access-Control-Allow-Credentials` HTTP response header.
	//
	// For more information about the `Access-Control-Allow-Credentials` HTTP response header, see [Access-Control-Allow-Credentials](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials) in the MDN Web Docs.
	AccessControlAllowCredentials interface{} `field:"required" json:"accessControlAllowCredentials" yaml:"accessControlAllowCredentials"`
	// A list of HTTP header names that CloudFront includes as values for the `Access-Control-Allow-Headers` HTTP response header.
	//
	// For more information about the `Access-Control-Allow-Headers` HTTP response header, see [Access-Control-Allow-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers) in the MDN Web Docs.
	AccessControlAllowHeaders interface{} `field:"required" json:"accessControlAllowHeaders" yaml:"accessControlAllowHeaders"`
	// A list of HTTP methods that CloudFront includes as values for the `Access-Control-Allow-Methods` HTTP response header.
	//
	// For more information about the `Access-Control-Allow-Methods` HTTP response header, see [Access-Control-Allow-Methods](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods) in the MDN Web Docs.
	AccessControlAllowMethods interface{} `field:"required" json:"accessControlAllowMethods" yaml:"accessControlAllowMethods"`
	// A list of origins (domain names) that CloudFront can use as the value for the `Access-Control-Allow-Origin` HTTP response header.
	//
	// For more information about the `Access-Control-Allow-Origin` HTTP response header, see [Access-Control-Allow-Origin](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin) in the MDN Web Docs.
	AccessControlAllowOrigins interface{} `field:"required" json:"accessControlAllowOrigins" yaml:"accessControlAllowOrigins"`
	// A Boolean that determines whether CloudFront overrides HTTP response headers received from the origin with the ones specified in this response headers policy.
	OriginOverride interface{} `field:"required" json:"originOverride" yaml:"originOverride"`
	// A list of HTTP headers that CloudFront includes as values for the `Access-Control-Expose-Headers` HTTP response header.
	//
	// For more information about the `Access-Control-Expose-Headers` HTTP response header, see [Access-Control-Expose-Headers](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers) in the MDN Web Docs.
	AccessControlExposeHeaders interface{} `field:"optional" json:"accessControlExposeHeaders" yaml:"accessControlExposeHeaders"`
	// A number that CloudFront uses as the value for the `Access-Control-Max-Age` HTTP response header.
	//
	// For more information about the `Access-Control-Max-Age` HTTP response header, see [Access-Control-Max-Age](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age) in the MDN Web Docs.
	AccessControlMaxAgeSec *float64 `field:"optional" json:"accessControlMaxAgeSec" yaml:"accessControlMaxAgeSec"`
}

