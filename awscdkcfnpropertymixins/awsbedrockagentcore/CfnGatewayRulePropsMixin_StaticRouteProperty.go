package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   staticRouteProperty := &StaticRouteProperty{
//   	Fallback: &FallbackProperty{
//   		CircuitBreaker: &CircuitBreakerConfigurationProperty{
//   			Mode: jsii.String("mode"),
//   		},
//   		FirstByteTimeoutMilliseconds: jsii.Number(123),
//   		TargetFallback: []interface{}{
//   			&FallbackTargetProperty{
//   				FirstByteTimeoutMilliseconds: jsii.Number(123),
//   				Model: jsii.String("model"),
//   				TargetName: jsii.String("targetName"),
//   				TriggerOn: &FallbackTriggersProperty{
//   					Errors: []*string{
//   						jsii.String("errors"),
//   					},
//   					StatusCodes: []*string{
//   						jsii.String("statusCodes"),
//   					},
//   				},
//   			},
//   		},
//   		TriggerOn: &FallbackTriggersProperty{
//   			Errors: []*string{
//   				jsii.String("errors"),
//   			},
//   			StatusCodes: []*string{
//   				jsii.String("statusCodes"),
//   			},
//   		},
//   	},
//   	Model: jsii.String("model"),
//   	TargetName: jsii.String("targetName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticroute.html
//
type CfnGatewayRulePropsMixin_StaticRouteProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticroute.html#cfn-bedrockagentcore-gatewayrule-staticroute-fallback
	//
	Fallback interface{} `field:"optional" json:"fallback" yaml:"fallback"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticroute.html#cfn-bedrockagentcore-gatewayrule-staticroute-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticroute.html#cfn-bedrockagentcore-gatewayrule-staticroute-targetname
	//
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
}

