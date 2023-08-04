package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   apigateway.NewRestApi(this, jsii.String("api"), &RestApiProps{
//   	DefaultCorsPreflightOptions: &CorsOptions{
//   		AllowOrigins: apigateway.Cors_ALL_ORIGINS(),
//   		AllowMethods: apigateway.Cors_ALL_METHODS(),
//   	},
//   })
//
type CorsOptions struct {
	// Specifies the list of origins that are allowed to make requests to this resource.
	//
	// If you wish to allow all origins, specify `Cors.ALL_ORIGINS` or
	// `[ * ]`.
	//
	// Responses will include the `Access-Control-Allow-Origin` response header.
	// If `Cors.ALL_ORIGINS` is specified, the `Vary: Origin` response header will
	// also be included.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
	//
	AllowOrigins *[]*string `field:"required" json:"allowOrigins" yaml:"allowOrigins"`
	// The Access-Control-Allow-Credentials response header tells browsers whether to expose the response to frontend JavaScript code when the request's credentials mode (Request.credentials) is "include".
	//
	// When a request's credentials mode (Request.credentials) is "include",
	// browsers will only expose the response to frontend JavaScript code if the
	// Access-Control-Allow-Credentials value is true.
	//
	// Credentials are cookies, authorization headers or TLS client certificates.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials
	//
	// Default: false.
	//
	AllowCredentials *bool `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// The Access-Control-Allow-Headers response header is used in response to a preflight request which includes the Access-Control-Request-Headers to indicate which HTTP headers can be used during the actual request.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers
	//
	// Default: Cors.DEFAULT_HEADERS
	//
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// The Access-Control-Allow-Methods response header specifies the method or methods allowed when accessing the resource in response to a preflight request.
	//
	// If `ANY` is specified, it will be expanded to `Cors.ALL_METHODS`.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods
	//
	// Default: Cors.ALL_METHODS
	//
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Sets Access-Control-Max-Age to -1, which means that caching is disabled.
	//
	// This option cannot be used with `maxAge`.
	// Default: - cache is enabled.
	//
	DisableCache *bool `field:"optional" json:"disableCache" yaml:"disableCache"`
	// The Access-Control-Expose-Headers response header indicates which headers can be exposed as part of the response by listing their names.
	//
	// If you want clients to be able to access other headers, you have to list
	// them using the Access-Control-Expose-Headers header.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers
	//
	// Default: - only the 6 CORS-safelisted response headers are exposed:
	// Cache-Control, Content-Language, Content-Type, Expires, Last-Modified,
	// Pragma.
	//
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// The Access-Control-Max-Age response header indicates how long the results of a preflight request (that is the information contained in the Access-Control-Allow-Methods and Access-Control-Allow-Headers headers) can be cached.
	//
	// To disable caching altogether use `disableCache: true`.
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age
	//
	// Default: - browser-specific (see reference).
	//
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Specifies the response status code returned from the OPTIONS method.
	// Default: 204.
	//
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
}

