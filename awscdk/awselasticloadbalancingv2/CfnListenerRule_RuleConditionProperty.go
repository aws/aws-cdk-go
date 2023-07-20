package awselasticloadbalancingv2


// Specifies a condition for a listener rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleConditionProperty := &RuleConditionProperty{
//   	Field: jsii.String("field"),
//   	HostHeaderConfig: &HostHeaderConfigProperty{
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	HttpHeaderConfig: &HttpHeaderConfigProperty{
//   		HttpHeaderName: jsii.String("httpHeaderName"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	HttpRequestMethodConfig: &HttpRequestMethodConfigProperty{
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	PathPatternConfig: &PathPatternConfigProperty{
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	QueryStringConfig: &QueryStringConfigProperty{
//   		Values: []interface{}{
//   			&QueryStringKeyValueProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	SourceIpConfig: &SourceIpConfigProperty{
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Information for a host header condition.
	//
	// Specify only when `Field` is `host-header` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-hostheaderconfig
	//
	HostHeaderConfig interface{} `field:"optional" json:"hostHeaderConfig" yaml:"hostHeaderConfig"`
	// Information for an HTTP header condition.
	//
	// Specify only when `Field` is `http-header` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-httpheaderconfig
	//
	HttpHeaderConfig interface{} `field:"optional" json:"httpHeaderConfig" yaml:"httpHeaderConfig"`
	// Information for an HTTP method condition.
	//
	// Specify only when `Field` is `http-request-method` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-httprequestmethodconfig
	//
	HttpRequestMethodConfig interface{} `field:"optional" json:"httpRequestMethodConfig" yaml:"httpRequestMethodConfig"`
	// Information for a path pattern condition.
	//
	// Specify only when `Field` is `path-pattern` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-pathpatternconfig
	//
	PathPatternConfig interface{} `field:"optional" json:"pathPatternConfig" yaml:"pathPatternConfig"`
	// Information for a query string condition.
	//
	// Specify only when `Field` is `query-string` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-querystringconfig
	//
	QueryStringConfig interface{} `field:"optional" json:"queryStringConfig" yaml:"queryStringConfig"`
	// Information for a source IP condition.
	//
	// Specify only when `Field` is `source-ip` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-sourceipconfig
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listenerrule-rulecondition.html#cfn-elasticloadbalancingv2-listenerrule-rulecondition-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

