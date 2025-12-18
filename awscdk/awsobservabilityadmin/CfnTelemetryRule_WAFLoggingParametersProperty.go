package awsobservabilityadmin


// Configuration parameters for WAF logging, including redacted fields and logging filters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wAFLoggingParametersProperty := &WAFLoggingParametersProperty{
//   	LoggingFilter: &LoggingFilterProperty{
//   		DefaultBehavior: jsii.String("defaultBehavior"),
//   		Filters: []interface{}{
//   			&FilterProperty{
//   				Behavior: jsii.String("behavior"),
//   				Conditions: []interface{}{
//   					&ConditionProperty{
//   						ActionCondition: &ActionConditionProperty{
//   							Action: jsii.String("action"),
//   						},
//   						LabelNameCondition: &LabelNameConditionProperty{
//   							LabelName: jsii.String("labelName"),
//   						},
//   					},
//   				},
//   				Requirement: jsii.String("requirement"),
//   			},
//   		},
//   	},
//   	LogType: jsii.String("logType"),
//   	RedactedFields: []interface{}{
//   		&FieldToMatchProperty{
//   			Method: jsii.String("method"),
//   			QueryString: jsii.String("queryString"),
//   			SingleHeader: &SingleHeaderProperty{
//   				Name: jsii.String("name"),
//   			},
//   			UriPath: jsii.String("uriPath"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-wafloggingparameters.html
//
type CfnTelemetryRule_WAFLoggingParametersProperty struct {
	// A filter configuration that determines which WAF log records to include or exclude.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-wafloggingparameters.html#cfn-observabilityadmin-telemetryrule-wafloggingparameters-loggingfilter
	//
	LoggingFilter interface{} `field:"optional" json:"loggingFilter" yaml:"loggingFilter"`
	// The type of WAF logs to collect (currently supports WAF_LOGS).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-wafloggingparameters.html#cfn-observabilityadmin-telemetryrule-wafloggingparameters-logtype
	//
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
	// The fields to redact from WAF logs to protect sensitive information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-wafloggingparameters.html#cfn-observabilityadmin-telemetryrule-wafloggingparameters-redactedfields
	//
	RedactedFields interface{} `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

