package awsbedrockagentcore


// Properties for CfnGatewayRateLimitPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnGatewayRateLimitMixinProps := &CfnGatewayRateLimitMixinProps{
//   	Description: jsii.String("description"),
//   	DimensionKeys: []*string{
//   		jsii.String("dimensionKeys"),
//   	},
//   	Entries: []interface{}{
//   		&LimitEntryProperty{
//   			Connections: []interface{}{
//   				&RateConfigProperty{
//   					Period: jsii.String("period"),
//   					Rate: jsii.Number(123),
//   				},
//   			},
//   			Dimensions: map[string]*string{
//   				"dimensionsKey": jsii.String("dimensions"),
//   			},
//   			Requests: []interface{}{
//   				&RateConfigProperty{
//   					Period: jsii.String("period"),
//   					Rate: jsii.Number(123),
//   				},
//   			},
//   			Tokens: []interface{}{
//   				&RateConfigProperty{
//   					Period: jsii.String("period"),
//   					Rate: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   	RateLimitId: jsii.String("rateLimitId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html
//
type CfnGatewayRateLimitMixinProps struct {
	// Optional human-readable description for this limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html#cfn-bedrockagentcore-gatewayratelimit-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Ordered list of dimension names defining the scope of this limit.
	//
	// Unique per gateway — no two limits can share the same dimensionKeys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html#cfn-bedrockagentcore-gatewayratelimit-dimensionkeys
	//
	DimensionKeys *[]*string `field:"optional" json:"dimensionKeys" yaml:"dimensionKeys"`
	// Rule entries mapping dimension values to rate configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html#cfn-bedrockagentcore-gatewayratelimit-entries
	//
	Entries interface{} `field:"optional" json:"entries" yaml:"entries"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html#cfn-bedrockagentcore-gatewayratelimit-gatewayidentifier
	//
	GatewayIdentifier *string `field:"optional" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
	// Limit identifier.
	//
	// Optional on Create (system-generates if not provided by customer).
	// Always present in responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html#cfn-bedrockagentcore-gatewayratelimit-ratelimitid
	//
	RateLimitId *string `field:"optional" json:"rateLimitId" yaml:"rateLimitId"`
}

