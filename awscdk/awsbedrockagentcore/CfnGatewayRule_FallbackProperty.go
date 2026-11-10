package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fallbackProperty := &FallbackProperty{
//   	TargetFallback: []interface{}{
//   		&FallbackTargetProperty{
//   			TargetName: jsii.String("targetName"),
//
//   			// the properties below are optional
//   			FirstByteTimeoutMilliseconds: jsii.Number(123),
//   			Model: jsii.String("model"),
//   			TriggerOn: &FallbackTriggersProperty{
//   				Errors: []*string{
//   					jsii.String("errors"),
//   				},
//   				StatusCodes: []*string{
//   					jsii.String("statusCodes"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	CircuitBreaker: &CircuitBreakerConfigurationProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   	FirstByteTimeoutMilliseconds: jsii.Number(123),
//   	TriggerOn: &FallbackTriggersProperty{
//   		Errors: []*string{
//   			jsii.String("errors"),
//   		},
//   		StatusCodes: []*string{
//   			jsii.String("statusCodes"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallback.html
//
type CfnGatewayRule_FallbackProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallback.html#cfn-bedrockagentcore-gatewayrule-fallback-targetfallback
	//
	TargetFallback interface{} `field:"required" json:"targetFallback" yaml:"targetFallback"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallback.html#cfn-bedrockagentcore-gatewayrule-fallback-circuitbreaker
	//
	CircuitBreaker interface{} `field:"optional" json:"circuitBreaker" yaml:"circuitBreaker"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallback.html#cfn-bedrockagentcore-gatewayrule-fallback-firstbytetimeoutmilliseconds
	//
	FirstByteTimeoutMilliseconds *float64 `field:"optional" json:"firstByteTimeoutMilliseconds" yaml:"firstByteTimeoutMilliseconds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-fallback.html#cfn-bedrockagentcore-gatewayrule-fallback-triggeron
	//
	TriggerOn interface{} `field:"optional" json:"triggerOn" yaml:"triggerOn"`
}

