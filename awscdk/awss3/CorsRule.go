package awss3


// Specifies a cross-origin access rule for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsRule := &CorsRule{
//   	AllowedMethods: []httpMethods{
//   		awscdk.Aws_s3.*httpMethods_GET,
//   	},
//   	AllowedOrigins: []*string{
//   		jsii.String("allowedOrigins"),
//   	},
//
//   	// the properties below are optional
//   	AllowedHeaders: []*string{
//   		jsii.String("allowedHeaders"),
//   	},
//   	ExposedHeaders: []*string{
//   		jsii.String("exposedHeaders"),
//   	},
//   	Id: jsii.String("id"),
//   	MaxAge: jsii.Number(123),
//   }
//
// Experimental.
type CorsRule struct {
	// An HTTP method that you allow the origin to execute.
	// Experimental.
	AllowedMethods *[]HttpMethods `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// One or more origins you want customers to be able to access the bucket from.
	// Experimental.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Headers that are specified in the Access-Control-Request-Headers header.
	// Experimental.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// One or more headers in the response that you want customers to be able to access from their applications.
	// Experimental.
	ExposedHeaders *[]*string `field:"optional" json:"exposedHeaders" yaml:"exposedHeaders"`
	// A unique identifier for this rule.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	// Experimental.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

