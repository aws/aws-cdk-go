package awswafv2


// Specifies a single custom aggregate key for a rate-base rule.
//
// > Web requests that are missing any of the components specified in the aggregation keys are omitted from the rate-based rule evaluation and handling.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var asn interface{}
//   var forwardedIp interface{}
//   var httpMethod interface{}
//   var ip interface{}
//
//   rateBasedStatementCustomKeyProperty := &RateBasedStatementCustomKeyProperty{
//   	Asn: asn,
//   	Cookie: &RateLimitCookieProperty{
//   		Name: jsii.String("name"),
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ForwardedIp: forwardedIp,
//   	Header: &RateLimitHeaderProperty{
//   		Name: jsii.String("name"),
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	HttpMethod: httpMethod,
//   	Ip: ip,
//   	Ja3Fingerprint: &RateLimitJA3FingerprintProperty{
//   		FallbackBehavior: jsii.String("fallbackBehavior"),
//   	},
//   	Ja4Fingerprint: &RateLimitJA4FingerprintProperty{
//   		FallbackBehavior: jsii.String("fallbackBehavior"),
//   	},
//   	LabelNamespace: &RateLimitLabelNamespaceProperty{
//   		Namespace: jsii.String("namespace"),
//   	},
//   	QueryArgument: &RateLimitQueryArgumentProperty{
//   		Name: jsii.String("name"),
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	QueryString: &RateLimitQueryStringProperty{
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	UriPath: &RateLimitUriPathProperty{
//   		TextTransformations: []interface{}{
//   			&TextTransformationProperty{
//   				Priority: jsii.Number(123),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html
//
type CfnWebACL_RateBasedStatementCustomKeyProperty struct {
	// Specifies the request's ASN as an aggregate key for a rate-based rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-asn
	//
	Asn interface{} `field:"optional" json:"asn" yaml:"asn"`
	// Use the value of a cookie in the request as an aggregate key.
	//
	// Each distinct value in the cookie contributes to the aggregation instance. If you use a single cookie as your custom key, then each value fully defines an aggregation instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-cookie
	//
	Cookie interface{} `field:"optional" json:"cookie" yaml:"cookie"`
	// Use the first IP address in an HTTP header as an aggregate key.
	//
	// Each distinct forwarded IP address contributes to the aggregation instance.
	//
	// When you specify an IP or forwarded IP in the custom key settings, you must also specify at least one other key to use. You can aggregate on only the forwarded IP address by specifying `FORWARDED_IP` in your rate-based statement's `AggregateKeyType` .
	//
	// With this option, you must specify the header to use in the rate-based rule's `ForwardedIPConfig` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-forwardedip
	//
	ForwardedIp interface{} `field:"optional" json:"forwardedIp" yaml:"forwardedIp"`
	// Use the value of a header in the request as an aggregate key.
	//
	// Each distinct value in the header contributes to the aggregation instance. If you use a single header as your custom key, then each value fully defines an aggregation instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-header
	//
	Header interface{} `field:"optional" json:"header" yaml:"header"`
	// Use the request's HTTP method as an aggregate key.
	//
	// Each distinct HTTP method contributes to the aggregation instance. If you use just the HTTP method as your custom key, then each method fully defines an aggregation instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-httpmethod
	//
	HttpMethod interface{} `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Use the request's originating IP address as an aggregate key. Each distinct IP address contributes to the aggregation instance.
	//
	// When you specify an IP or forwarded IP in the custom key settings, you must also specify at least one other key to use. You can aggregate on only the IP address by specifying `IP` in your rate-based statement's `AggregateKeyType` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-ip
	//
	Ip interface{} `field:"optional" json:"ip" yaml:"ip"`
	// Use the request's JA3 fingerprint as an aggregate key.
	//
	// If you use a single JA3 fingerprint as your custom key, then each value fully defines an aggregation instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-ja3fingerprint
	//
	Ja3Fingerprint interface{} `field:"optional" json:"ja3Fingerprint" yaml:"ja3Fingerprint"`
	// Use the request's JA4 fingerprint as an aggregate key.
	//
	// If you use a single JA4 fingerprint as your custom key, then each value fully defines an aggregation instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-ja4fingerprint
	//
	Ja4Fingerprint interface{} `field:"optional" json:"ja4Fingerprint" yaml:"ja4Fingerprint"`
	// Use the specified label namespace as an aggregate key.
	//
	// Each distinct fully qualified label name that has the specified label namespace contributes to the aggregation instance. If you use just one label namespace as your custom key, then each label name fully defines an aggregation instance.
	//
	// This uses only labels that have been added to the request by rules that are evaluated before this rate-based rule in the web ACL.
	//
	// For information about label namespaces and names, see [Label syntax and naming requirements](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-label-requirements.html) in the *AWS WAF Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-labelnamespace
	//
	LabelNamespace interface{} `field:"optional" json:"labelNamespace" yaml:"labelNamespace"`
	// Use the specified query argument as an aggregate key.
	//
	// Each distinct value for the named query argument contributes to the aggregation instance. If you use a single query argument as your custom key, then each value fully defines an aggregation instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-queryargument
	//
	QueryArgument interface{} `field:"optional" json:"queryArgument" yaml:"queryArgument"`
	// Use the request's query string as an aggregate key.
	//
	// Each distinct string contributes to the aggregation instance. If you use just the query string as your custom key, then each string fully defines an aggregation instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-querystring
	//
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// Use the request's URI path as an aggregate key.
	//
	// Each distinct URI path contributes to the aggregation instance. If you use just the URI path as your custom key, then each URI path fully defines an aggregation instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatementcustomkey.html#cfn-wafv2-webacl-ratebasedstatementcustomkey-uripath
	//
	UriPath interface{} `field:"optional" json:"uriPath" yaml:"uriPath"`
}

