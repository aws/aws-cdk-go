package awsobservabilityadmin


// Configuration specifying where and how telemetry data should be delivered for AWS resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telemetryDestinationConfigurationProperty := &TelemetryDestinationConfigurationProperty{
//   	CloudtrailParameters: &CloudtrailParametersProperty{
//   		AdvancedEventSelectors: []interface{}{
//   			&AdvancedEventSelectorProperty{
//   				FieldSelectors: []interface{}{
//   					&AdvancedFieldSelectorProperty{
//   						EndsWith: []*string{
//   							jsii.String("endsWith"),
//   						},
//   						EqualTo: []*string{
//   							jsii.String("equalTo"),
//   						},
//   						Field: jsii.String("field"),
//   						NotEndsWith: []*string{
//   							jsii.String("notEndsWith"),
//   						},
//   						NotEquals: []*string{
//   							jsii.String("notEquals"),
//   						},
//   						NotStartsWith: []*string{
//   							jsii.String("notStartsWith"),
//   						},
//   						StartsWith: []*string{
//   							jsii.String("startsWith"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	DestinationPattern: jsii.String("destinationPattern"),
//   	DestinationType: jsii.String("destinationType"),
//   	ElbLoadBalancerLoggingParameters: &ELBLoadBalancerLoggingParametersProperty{
//   		FieldDelimiter: jsii.String("fieldDelimiter"),
//   		OutputFormat: jsii.String("outputFormat"),
//   	},
//   	LogDeliveryParameters: &LogDeliveryParametersProperty{
//   		LogTypes: []*string{
//   			jsii.String("logTypes"),
//   		},
//   	},
//   	RetentionInDays: jsii.Number(123),
//   	VpcFlowLogParameters: &VPCFlowLogParametersProperty{
//   		LogFormat: jsii.String("logFormat"),
//   		MaxAggregationInterval: jsii.Number(123),
//   		TrafficType: jsii.String("trafficType"),
//   	},
//   	WafLoggingParameters: &WAFLoggingParametersProperty{
//   		LoggingFilter: &LoggingFilterProperty{
//   			DefaultBehavior: jsii.String("defaultBehavior"),
//   			Filters: []interface{}{
//   				&FilterProperty{
//   					Behavior: jsii.String("behavior"),
//   					Conditions: []interface{}{
//   						&ConditionProperty{
//   							ActionCondition: &ActionConditionProperty{
//   								Action: jsii.String("action"),
//   							},
//   							LabelNameCondition: &LabelNameConditionProperty{
//   								LabelName: jsii.String("labelName"),
//   							},
//   						},
//   					},
//   					Requirement: jsii.String("requirement"),
//   				},
//   			},
//   		},
//   		LogType: jsii.String("logType"),
//   		RedactedFields: []interface{}{
//   			&FieldToMatchProperty{
//   				Method: jsii.String("method"),
//   				QueryString: jsii.String("queryString"),
//   				SingleHeader: &SingleHeaderProperty{
//   					Name: jsii.String("name"),
//   				},
//   				UriPath: jsii.String("uriPath"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-telemetrydestinationconfiguration.html
//
type CfnTelemetryRule_TelemetryDestinationConfigurationProperty struct {
	// Configuration parameters specific to AWS CloudTrail when CloudTrail is the source type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-telemetryrule-telemetrydestinationconfiguration-cloudtrailparameters
	//
	CloudtrailParameters interface{} `field:"optional" json:"cloudtrailParameters" yaml:"cloudtrailParameters"`
	// The pattern used to generate the destination path or name, supporting macros like <resourceId> and <accountId>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-telemetryrule-telemetrydestinationconfiguration-destinationpattern
	//
	DestinationPattern *string `field:"optional" json:"destinationPattern" yaml:"destinationPattern"`
	// The type of destination for the telemetry data (e.g., "Amazon CloudWatch Logs", "S3").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-telemetryrule-telemetrydestinationconfiguration-destinationtype
	//
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Configuration parameters specific to ELB load balancer logging when ELB is the resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-telemetryrule-telemetrydestinationconfiguration-elbloadbalancerloggingparameters
	//
	ElbLoadBalancerLoggingParameters interface{} `field:"optional" json:"elbLoadBalancerLoggingParameters" yaml:"elbLoadBalancerLoggingParameters"`
	// Configuration parameters specific to Amazon Bedrock AgentCore logging when Amazon Bedrock AgentCore is the resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-telemetryrule-telemetrydestinationconfiguration-logdeliveryparameters
	//
	LogDeliveryParameters interface{} `field:"optional" json:"logDeliveryParameters" yaml:"logDeliveryParameters"`
	// The number of days to retain the telemetry data in the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-telemetryrule-telemetrydestinationconfiguration-retentionindays
	//
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Configuration parameters specific to VPC Flow Logs when VPC is the resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-telemetryrule-telemetrydestinationconfiguration-vpcflowlogparameters
	//
	VpcFlowLogParameters interface{} `field:"optional" json:"vpcFlowLogParameters" yaml:"vpcFlowLogParameters"`
	// Configuration parameters specific to WAF logging when WAF is the resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-telemetryrule-telemetrydestinationconfiguration.html#cfn-observabilityadmin-telemetryrule-telemetrydestinationconfiguration-wafloggingparameters
	//
	WafLoggingParameters interface{} `field:"optional" json:"wafLoggingParameters" yaml:"wafLoggingParameters"`
}

