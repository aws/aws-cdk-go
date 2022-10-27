package awss3


// All http request methods.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &bucketProps{
//   	websiteRoutingRules: []routingRule{
//   		&routingRule{
//   			hostName: jsii.String("www.example.com"),
//   			httpRedirectCode: jsii.String("302"),
//   			protocol: s3.redirectProtocol_HTTPS,
//   			replaceKey: s3.replaceKey.prefixWith(jsii.String("test/")),
//   			condition: &routingRuleCondition{
//   				httpErrorCodeReturnedEquals: jsii.String("200"),
//   				keyPrefixEquals: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   })
//
type RedirectProtocol string

const (
	RedirectProtocol_HTTP RedirectProtocol = "HTTP"
	RedirectProtocol_HTTPS RedirectProtocol = "HTTPS"
)

