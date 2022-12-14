package awswafv2


// Properties for defining a `CfnLoggingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var method interface{}
//   var queryString interface{}
//   var uriPath interface{}
//
//   cfnLoggingConfigurationProps := &cfnLoggingConfigurationProps{
//   	logDestinationConfigs: []*string{
//   		jsii.String("logDestinationConfigs"),
//   	},
//   	resourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	loggingFilter: &loggingFilterProperty{
//   		defaultBehavior: jsii.String("defaultBehavior"),
//   		filters: []interface{}{
//   			&filterProperty{
//   				behavior: jsii.String("behavior"),
//   				conditions: []interface{}{
//   					&conditionProperty{
//   						actionCondition: &actionConditionProperty{
//   							action: jsii.String("action"),
//   						},
//   						labelNameCondition: &labelNameConditionProperty{
//   							labelName: jsii.String("labelName"),
//   						},
//   					},
//   				},
//   				requirement: jsii.String("requirement"),
//   			},
//   		},
//   	},
//   	redactedFields: []interface{}{
//   		&fieldToMatchProperty{
//   			jsonBody: &jsonBodyProperty{
//   				matchPattern: &matchPatternProperty{
//   					all: all,
//   					includedPaths: []*string{
//   						jsii.String("includedPaths"),
//   					},
//   				},
//   				matchScope: jsii.String("matchScope"),
//
//   				// the properties below are optional
//   				invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   			},
//   			method: method,
//   			queryString: queryString,
//   			singleHeader: &singleHeaderProperty{
//   				name: jsii.String("name"),
//   			},
//   			uriPath: uriPath,
//   		},
//   	},
//   }
//
type CfnLoggingConfigurationProps struct {
	// The logging destination configuration that you want to associate with the web ACL.
	//
	// > You can associate one logging destination to a web ACL.
	LogDestinationConfigs *[]*string `field:"required" json:"logDestinationConfigs" yaml:"logDestinationConfigs"`
	// The Amazon Resource Name (ARN) of the web ACL that you want to associate with `LogDestinationConfigs` .
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// Filtering that specifies which web requests are kept in the logs and which are dropped.
	//
	// You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.
	LoggingFilter interface{} `field:"optional" json:"loggingFilter" yaml:"loggingFilter"`
	// The parts of the request that you want to keep out of the logs.
	//
	// For example, if you redact the `SingleHeader` field, the `HEADER` field in the logs will be `xxx` .
	//
	// > You can specify only the following fields for redaction: `UriPath` , `QueryString` , `SingleHeader` , `Method` , and `JsonBody` .
	RedactedFields interface{} `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

