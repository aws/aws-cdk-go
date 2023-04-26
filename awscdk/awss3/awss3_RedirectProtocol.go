package awss3


// All http request methods.
//
// Example:
//   bucket := s3.NewBucket(this, jsii.String("MyRedirectedBucket"), &BucketProps{
//   	WebsiteRoutingRules: []routingRule{
//   		&routingRule{
//   			HostName: jsii.String("www.example.com"),
//   			HttpRedirectCode: jsii.String("302"),
//   			Protocol: s3.RedirectProtocol_HTTPS,
//   			ReplaceKey: s3.ReplaceKey_PrefixWith(jsii.String("test/")),
//   			Condition: &RoutingRuleCondition{
//   				HttpErrorCodeReturnedEquals: jsii.String("200"),
//   				KeyPrefixEquals: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type RedirectProtocol string

const (
	// Experimental.
	RedirectProtocol_HTTP RedirectProtocol = "HTTP"
	// Experimental.
	RedirectProtocol_HTTPS RedirectProtocol = "HTTPS"
)

