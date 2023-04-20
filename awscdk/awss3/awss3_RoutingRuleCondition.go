package awss3


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
type RoutingRuleCondition struct {
	// The HTTP error code when the redirect is applied.
	//
	// In the event of an error, if the error code equals this value, then the specified redirect is applied.
	//
	// If both condition properties are specified, both must be true for the redirect to be applied.
	// Experimental.
	HttpErrorCodeReturnedEquals *string `field:"optional" json:"httpErrorCodeReturnedEquals" yaml:"httpErrorCodeReturnedEquals"`
	// The object key name prefix when the redirect is applied.
	//
	// If both condition properties are specified, both must be true for the redirect to be applied.
	// Experimental.
	KeyPrefixEquals *string `field:"optional" json:"keyPrefixEquals" yaml:"keyPrefixEquals"`
}

