package awselasticloadbalancingv2


// A redirect response.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redirectResponse := &redirectResponse{
//   	statusCode: jsii.String("statusCode"),
//
//   	// the properties below are optional
//   	host: jsii.String("host"),
//   	path: jsii.String("path"),
//   	port: jsii.String("port"),
//   	protocol: jsii.String("protocol"),
//   	query: jsii.String("query"),
//   }
//
// Deprecated: superceded by `ListenerAction.redirect()`.
type RedirectResponse struct {
	// The HTTP redirect code (HTTP_301 or HTTP_302).
	// Deprecated: superceded by `ListenerAction.redirect()`.
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The hostname.
	//
	// This component is not percent-encoded. The hostname can contain #{host}.
	// Deprecated: superceded by `ListenerAction.redirect()`.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The absolute path, starting with the leading "/".
	//
	// This component is not percent-encoded.
	// The path can contain #{host}, #{path}, and #{port}.
	// Deprecated: superceded by `ListenerAction.redirect()`.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port.
	//
	// You can specify a value from 1 to 65535 or #{port}.
	// Deprecated: superceded by `ListenerAction.redirect()`.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol.
	//
	// You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP,
	// HTTP to HTTPS, and HTTPS to HTTPS. You cannot redirect HTTPS to HTTP.
	// Deprecated: superceded by `ListenerAction.redirect()`.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded.
	//
	// Do not include the leading "?", as it is automatically added.
	// You can specify any of the reserved keywords.
	// Deprecated: superceded by `ListenerAction.redirect()`.
	Query *string `field:"optional" json:"query" yaml:"query"`
}

