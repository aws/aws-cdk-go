package awscloudfront


// A list of HTTP methods that CloudFront includes as values for the `Access-Control-Allow-Methods` HTTP response header.
//
// For more information about the `Access-Control-Allow-Methods` HTTP response header, see [Access-Control-Allow-Methods](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods) in the MDN Web Docs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessControlAllowMethodsProperty := &accessControlAllowMethodsProperty{
//   	items: []*string{
//   		jsii.String("items"),
//   	},
//   }
//
type CfnResponseHeadersPolicy_AccessControlAllowMethodsProperty struct {
	// The list of HTTP methods. Valid values are:.
	//
	// - `GET`
	// - `DELETE`
	// - `HEAD`
	// - `OPTIONS`
	// - `PATCH`
	// - `POST`
	// - `PUT`
	// - `ALL`
	//
	// `ALL` is a special value that includes all of the listed HTTP methods.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
}

