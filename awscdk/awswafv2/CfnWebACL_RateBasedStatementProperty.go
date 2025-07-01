package awswafv2


// A rate-based rule counts incoming requests and rate limits requests when they are coming at too fast a rate.
//
// The rule categorizes requests according to your aggregation criteria, collects them into aggregation instances, and counts and rate limits the requests for each instance.
//
// > If you change any of these settings in a rule that's currently in use, the change resets the rule's rate limiting counts. This can pause the rule's rate limiting activities for up to a minute.
//
// You can specify individual aggregation keys, like IP address or HTTP method. You can also specify aggregation key combinations, like IP address and HTTP method, or HTTP method, query argument, and cookie.
//
// Each unique set of values for the aggregation keys that you specify is a separate aggregation instance, with the value from each key contributing to the aggregation instance definition.
//
// For example, assume the rule evaluates web requests with the following IP address and HTTP method values:
//
// - IP address 10.1.1.1, HTTP method POST
// - IP address 10.1.1.1, HTTP method GET
// - IP address 127.0.0.0, HTTP method POST
// - IP address 10.1.1.1, HTTP method GET
//
// The rule would create different aggregation instances according to your aggregation criteria, for example:
//
// - If the aggregation criteria is just the IP address, then each individual address is an aggregation instance, and AWS WAF counts requests separately for each. The aggregation instances and request counts for our example would be the following:
//
// - IP address 10.1.1.1: count 3
// - IP address 127.0.0.0: count 1
// - If the aggregation criteria is HTTP method, then each individual HTTP method is an aggregation instance. The aggregation instances and request counts for our example would be the following:
//
// - HTTP method POST: count 2
// - HTTP method GET: count 2
// - If the aggregation criteria is IP address and HTTP method, then each IP address and each HTTP method would contribute to the combined aggregation instance. The aggregation instances and request counts for our example would be the following:
//
// - IP address 10.1.1.1, HTTP method POST: count 1
// - IP address 10.1.1.1, HTTP method GET: count 2
// - IP address 127.0.0.0, HTTP method POST: count 1
//
// For any n-tuple of aggregation keys, each unique combination of values for the keys defines a separate aggregation instance, which AWS WAF counts and rate-limits individually.
//
// You can optionally nest another statement inside the rate-based statement, to narrow the scope of the rule so that it only counts and rate limits requests that match the nested statement. You can use this nested scope-down statement in conjunction with your aggregation key specifications or you can just count and rate limit all requests that match the scope-down statement, without additional aggregation. When you choose to just manage all requests that match a scope-down statement, the aggregation instance is singular for the rule.
//
// You cannot nest a `RateBasedStatement` inside another statement, for example inside a `NotStatement` or `OrStatement` . You can define a `RateBasedStatement` inside a web ACL and inside a rule group.
//
// For additional information about the options, see [Rate limiting web requests using rate-based rules](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rate-based-rules.html) in the *AWS WAF Developer Guide* .
//
// If you only aggregate on the individual IP address or forwarded IP address, you can retrieve the list of IP addresses that AWS WAF is currently rate limiting for a rule through the API call `GetRateBasedStatementManagedKeys` . This option is not available for other aggregation configurations.
//
// AWS WAF tracks and manages web requests separately for each instance of a rate-based rule that you use. For example, if you provide the same rate-based rule settings in two web ACLs, each of the two rule statements represents a separate instance of the rate-based rule and gets its own tracking and management by AWS WAF . If you define a rate-based rule inside a rule group, and then use that rule group in multiple places, each use creates a separate instance of the rate-based rule that gets its own tracking and management by AWS WAF .
//
// Example:
//
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html
//
type CfnWebACL_RateBasedStatementProperty struct {
	// Setting that indicates how to aggregate the request counts.
	//
	// > Web requests that are missing any of the components specified in the aggregation keys are omitted from the rate-based rule evaluation and handling.
	//
	// - `CONSTANT` - Count and limit the requests that match the rate-based rule's scope-down statement. With this option, the counted requests aren't further aggregated. The scope-down statement is the only specification used. When the count of all requests that satisfy the scope-down statement goes over the limit, AWS WAF applies the rule action to all requests that satisfy the scope-down statement.
	//
	// With this option, you must configure the `ScopeDownStatement` property.
	// - `CUSTOM_KEYS` - Aggregate the request counts using one or more web request components as the aggregate keys.
	//
	// With this option, you must specify the aggregate keys in the `CustomKeys` property.
	//
	// To aggregate on only the IP address or only the forwarded IP address, don't use custom keys. Instead, set the aggregate key type to `IP` or `FORWARDED_IP` .
	// - `FORWARDED_IP` - Aggregate the request counts on the first IP address in an HTTP header.
	//
	// With this option, you must specify the header to use in the `ForwardedIPConfig` property.
	//
	// To aggregate on a combination of the forwarded IP address with other aggregate keys, use `CUSTOM_KEYS` .
	// - `IP` - Aggregate the request counts on the IP address from the web request origin.
	//
	// To aggregate on a combination of the IP address with other aggregate keys, use `CUSTOM_KEYS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-aggregatekeytype
	//
	AggregateKeyType *string `field:"required" json:"aggregateKeyType" yaml:"aggregateKeyType"`
	// The limit on requests during the specified evaluation window for a single aggregation instance for the rate-based rule.
	//
	// If the rate-based statement includes a `ScopeDownStatement` , this limit is applied only to the requests that match the statement.
	//
	// Examples:
	//
	// - If you aggregate on just the IP address, this is the limit on requests from any single IP address.
	// - If you aggregate on the HTTP method and the query argument name "city", then this is the limit on requests for any single method, city pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-limit
	//
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
	// Specifies the aggregate keys to use in a rate-base rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-customkeys
	//
	CustomKeys interface{} `field:"optional" json:"customKeys" yaml:"customKeys"`
	// The amount of time, in seconds, that AWS WAF should include in its request counts, looking back from the current time.
	//
	// For example, for a setting of 120, when AWS WAF checks the rate, it counts the requests for the 2 minutes immediately preceding the current time. Valid settings are 60, 120, 300, and 600.
	//
	// This setting doesn't determine how often AWS WAF checks the rate, but how far back it looks each time it checks. AWS WAF checks the rate about every 10 seconds.
	//
	// Default: `300` (5 minutes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-evaluationwindowsec
	//
	EvaluationWindowSec *float64 `field:"optional" json:"evaluationWindowSec" yaml:"evaluationWindowSec"`
	// The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin.
	//
	// Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.
	//
	// > If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.
	//
	// This is required if you specify a forwarded IP in the rule's aggregate key settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-forwardedipconfig
	//
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
	// An optional nested statement that narrows the scope of the web requests that are evaluated and managed by the rate-based statement.
	//
	// When you use a scope-down statement, the rate-based rule only tracks and rate limits requests that match the scope-down statement. You can use any nestable `Statement` in the scope-down statement, and you can nest statements at any level, the same as you can for a rule statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-webacl-ratebasedstatement.html#cfn-wafv2-webacl-ratebasedstatement-scopedownstatement
	//
	ScopeDownStatement interface{} `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
}

