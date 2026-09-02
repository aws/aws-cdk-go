package awsbedrockagentcore


// A single rule entry within a limit, mapping dimension values to rate configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   limitEntryProperty := &LimitEntryProperty{
//   	Connections: []interface{}{
//   		&RateConfigProperty{
//   			Period: jsii.String("period"),
//   			Rate: jsii.Number(123),
//   		},
//   	},
//   	Dimensions: map[string]*string{
//   		"dimensionsKey": jsii.String("dimensions"),
//   	},
//   	Requests: []interface{}{
//   		&RateConfigProperty{
//   			Period: jsii.String("period"),
//   			Rate: jsii.Number(123),
//   		},
//   	},
//   	Tokens: []interface{}{
//   		&RateConfigProperty{
//   			Period: jsii.String("period"),
//   			Rate: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayratelimit-limitentry.html
//
type CfnGatewayRateLimitPropsMixin_LimitEntryProperty struct {
	// Connection rate limits (per second only).
	//
	// Limited to 1 entry for now. — P2
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayratelimit-limitentry.html#cfn-bedrockagentcore-gatewayratelimit-limitentry-connections
	//
	Connections interface{} `field:"optional" json:"connections" yaml:"connections"`
	// Map of dimension name to dimension value for a rule entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayratelimit-limitentry.html#cfn-bedrockagentcore-gatewayratelimit-limitentry-dimensions
	//
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Request rate limits (RPS or RPM).
	//
	// Limited to 1 entry for now.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayratelimit-limitentry.html#cfn-bedrockagentcore-gatewayratelimit-limitentry-requests
	//
	Requests interface{} `field:"optional" json:"requests" yaml:"requests"`
	// Token rate limits (TPM).
	//
	// Limited to 1 entry for now. — P1
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayratelimit-limitentry.html#cfn-bedrockagentcore-gatewayratelimit-limitentry-tokens
	//
	Tokens interface{} `field:"optional" json:"tokens" yaml:"tokens"`
}

