package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   targetTrafficSplitEntryProperty := &TargetTrafficSplitEntryProperty{
//   	Description: jsii.String("description"),
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
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   	Model: jsii.String("model"),
//   	Name: jsii.String("name"),
//   	TargetName: jsii.String("targetName"),
//   	Weight: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-targettrafficsplitentry.html
//
type CfnGatewayRulePropsMixin_TargetTrafficSplitEntryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-targettrafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-targettrafficsplitentry-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-targettrafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-targettrafficsplitentry-fallback
	//
	Fallback interface{} `field:"optional" json:"fallback" yaml:"fallback"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-targettrafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-targettrafficsplitentry-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-targettrafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-targettrafficsplitentry-model
	//
	Model *string `field:"optional" json:"model" yaml:"model"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-targettrafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-targettrafficsplitentry-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-targettrafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-targettrafficsplitentry-targetname
	//
	TargetName *string `field:"optional" json:"targetName" yaml:"targetName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-targettrafficsplitentry.html#cfn-bedrockagentcore-gatewayrule-targettrafficsplitentry-weight
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

