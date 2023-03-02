package awsmediastore


// A rule for a CORS policy.
//
// You can add up to 100 rules to a CORS policy. If more than one rule applies, the service uses the first applicable rule listed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsRuleProperty := &corsRuleProperty{
//   	allowedHeaders: []*string{
//   		jsii.String("allowedHeaders"),
//   	},
//   	allowedMethods: []*string{
//   		jsii.String("allowedMethods"),
//   	},
//   	allowedOrigins: []*string{
//   		jsii.String("allowedOrigins"),
//   	},
//   	exposeHeaders: []*string{
//   		jsii.String("exposeHeaders"),
//   	},
//   	maxAgeSeconds: jsii.Number(123),
//   }
//
type CfnContainer_CorsRuleProperty struct {
	// Specifies which headers are allowed in a preflight `OPTIONS` request through the `Access-Control-Request-Headers` header.
	//
	// Each header name that is specified in `Access-Control-Request-Headers` must have a corresponding entry in the rule. Only the headers that were requested are sent back.
	//
	// This element can contain only one wildcard character (*).
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Identifies an HTTP method that the origin that is specified in the rule is allowed to execute.
	//
	// Each CORS rule must contain at least one `AllowedMethods` and one `AllowedOrigins` element.
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// One or more response headers that you want users to be able to access from their applications (for example, from a JavaScript `XMLHttpRequest` object).
	//
	// Each CORS rule must have at least one `AllowedOrigins` element. The string value can include only one wildcard character (*), for example, http://*.example.com. Additionally, you can specify only one wildcard character to allow cross-origin access for all origins.
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// One or more headers in the response that you want users to be able to access from their applications (for example, from a JavaScript `XMLHttpRequest` object).
	//
	// This element is optional for each rule.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// The time in seconds that your browser caches the preflight response for the specified resource.
	//
	// A CORS rule can have only one `MaxAgeSeconds` element.
	MaxAgeSeconds *float64 `field:"optional" json:"maxAgeSeconds" yaml:"maxAgeSeconds"`
}

