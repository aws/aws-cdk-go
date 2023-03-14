package awselasticloadbalancingv2


// Specifies a condition for a listener rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleConditionProperty := &ruleConditionProperty{
//   	field: jsii.String("field"),
//   	hostHeaderConfig: &hostHeaderConfigProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	httpHeaderConfig: &httpHeaderConfigProperty{
//   		httpHeaderName: jsii.String("httpHeaderName"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	httpRequestMethodConfig: &httpRequestMethodConfigProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	pathPatternConfig: &pathPatternConfigProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	queryStringConfig: &queryStringConfigProperty{
//   		values: []interface{}{
//   			&queryStringKeyValueProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	sourceIpConfig: &sourceIpConfigProperty{
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnListenerRule_RuleConditionProperty struct {
	// The field in the HTTP request. The following are the possible values:.
	//
	// - `http-header`
	// - `http-request-method`
	// - `host-header`
	// - `path-pattern`
	// - `query-string`
	// - `source-ip`.
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Information for a host header condition.
	//
	// Specify only when `Field` is `host-header` .
	HostHeaderConfig interface{} `field:"optional" json:"hostHeaderConfig" yaml:"hostHeaderConfig"`
	// Information for an HTTP header condition.
	//
	// Specify only when `Field` is `http-header` .
	HttpHeaderConfig interface{} `field:"optional" json:"httpHeaderConfig" yaml:"httpHeaderConfig"`
	// Information for an HTTP method condition.
	//
	// Specify only when `Field` is `http-request-method` .
	HttpRequestMethodConfig interface{} `field:"optional" json:"httpRequestMethodConfig" yaml:"httpRequestMethodConfig"`
	// Information for a path pattern condition.
	//
	// Specify only when `Field` is `path-pattern` .
	PathPatternConfig interface{} `field:"optional" json:"pathPatternConfig" yaml:"pathPatternConfig"`
	// Information for a query string condition.
	//
	// Specify only when `Field` is `query-string` .
	QueryStringConfig interface{} `field:"optional" json:"queryStringConfig" yaml:"queryStringConfig"`
	// Information for a source IP condition.
	//
	// Specify only when `Field` is `source-ip` .
	SourceIpConfig interface{} `field:"optional" json:"sourceIpConfig" yaml:"sourceIpConfig"`
	// The condition value.
	//
	// Specify only when `Field` is `host-header` or `path-pattern` . Alternatively, to specify multiple host names or multiple path patterns, use `HostHeaderConfig` or `PathPatternConfig` .
	//
	// If `Field` is `host-header` and you're not using `HostHeaderConfig` , you can specify a single host name (for example, my.example.com). A host name is case insensitive, can be up to 128 characters in length, and can contain any of the following characters.
	//
	// - A-Z, a-z, 0-9
	// - - .
	// - * (matches 0 or more characters)
	// - ? (matches exactly 1 character)
	//
	// If `Field` is `path-pattern` and you're not using `PathPatternConfig` , you can specify a single path pattern (for example, /img/*). A path pattern is case-sensitive, can be up to 128 characters in length, and can contain any of the following characters.
	//
	// - A-Z, a-z, 0-9
	// - _ - . $ / ~ " ' @ : +
	// - & (using &amp;)
	// - * (matches 0 or more characters)
	// - ? (matches exactly 1 character)
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

