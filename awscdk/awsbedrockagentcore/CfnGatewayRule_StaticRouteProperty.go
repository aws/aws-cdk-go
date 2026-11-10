package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   staticRouteProperty := &StaticRouteProperty{
//   	TargetName: jsii.String("targetName"),
//
//   	// the properties below are optional
//   	Fallback: &FallbackProperty{
//   		TargetFallback: []interface{}{
//   			&FallbackTargetProperty{
//   				TargetName: jsii.String("targetName"),
//
//   				// the properties below are optional
//   				FirstByteTimeoutMilliseconds: jsii.Number(123),
//   				Model: jsii.String("model"),
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
//
//   		// the properties below are optional
//   		CircuitBreaker: &CircuitBreakerConfigurationProperty{
//   			Mode: jsii.String("mode"),
//   		},
//   		FirstByteTimeoutMilliseconds: jsii.Number(123),
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticroute.html
//
type CfnGatewayRule_StaticRouteProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticroute.html#cfn-bedrockagentcore-gatewayrule-staticroute-targetname
	//
	TargetName *string `field:"required" json:"targetName" yaml:"targetName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticroute.html#cfn-bedrockagentcore-gatewayrule-staticroute-fallback
	//
	Fallback interface{} `field:"optional" json:"fallback" yaml:"fallback"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-staticroute.html#cfn-bedrockagentcore-gatewayrule-staticroute-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
}

