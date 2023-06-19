package awswafv2


// Properties for defining a `CfnLoggingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var jsonBody interface{}
//   var loggingFilter interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var uriPath interface{}
//
//   cfnLoggingConfigurationProps := &CfnLoggingConfigurationProps{
//   	LogDestinationConfigs: []*string{
//   		jsii.String("logDestinationConfigs"),
//   	},
//   	ResourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	LoggingFilter: loggingFilter,
//   	RedactedFields: []interface{}{
//   		&FieldToMatchProperty{
//   			JsonBody: jsonBody,
//   			Method: method,
//   			QueryString: queryString,
//   			SingleHeader: singleHeader,
//   			UriPath: uriPath,
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
	// For example, if you redact the `SingleHeader` field, the `HEADER` field in the logs will be `REDACTED` for all rules that use the `SingleHeader` `FieldToMatch` setting.
	//
	// Redaction applies only to the component that's specified in the rule's `FieldToMatch` setting, so the `SingleHeader` redaction doesn't apply to rules that use the `Headers` `FieldToMatch` .
	//
	// > You can specify only the following fields for redaction: `UriPath` , `QueryString` , `SingleHeader` , and `Method` .
	RedactedFields interface{} `field:"optional" json:"redactedFields" yaml:"redactedFields"`
}

