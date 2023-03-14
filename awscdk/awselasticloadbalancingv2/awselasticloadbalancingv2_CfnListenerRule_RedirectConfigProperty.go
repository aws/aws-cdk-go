package awselasticloadbalancingv2


// Information about a redirect action.
//
// A URI consists of the following components: protocol://hostname:port/path?query. You must modify at least one of the following components to avoid a redirect loop: protocol, hostname, port, or path. Any components that you do not modify retain their original values.
//
// You can reuse URI components using the following reserved keywords:
//
// - #{protocol}
// - #{host}
// - #{port}
// - #{path} (the leading "/" is removed)
// - #{query}
//
// For example, you can change the path to "/new/#{path}", the hostname to "example.#{host}", or the query to "#{query}&value=xyz".
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redirectConfigProperty := &redirectConfigProperty{
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
type CfnListenerRule_RedirectConfigProperty struct {
	// The HTTP redirect code.
	//
	// The redirect is either permanent (HTTP 301) or temporary (HTTP 302).
	StatusCode *string `field:"required" json:"statusCode" yaml:"statusCode"`
	// The hostname.
	//
	// This component is not percent-encoded. The hostname can contain #{host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The absolute path, starting with the leading "/".
	//
	// This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The port.
	//
	// You can specify a value from 1 to 65535 or #{port}.
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocol.
	//
	// You can specify HTTP, HTTPS, or #{protocol}. You can redirect HTTP to HTTP, HTTP to HTTPS, and HTTPS to HTTPS. You cannot redirect HTTPS to HTTP.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The query parameters, URL-encoded when necessary, but not percent-encoded.
	//
	// Do not include the leading "?", as it is automatically added. You can specify any of the reserved keywords.
	Query *string `field:"optional" json:"query" yaml:"query"`
}

