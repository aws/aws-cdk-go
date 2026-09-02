package awsbedrockagentcore


// Properties for defining a `CfnGatewayRateLimit`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayRateLimitProps := &CfnGatewayRateLimitProps{
//   	DimensionKeys: []*string{
//   		jsii.String("dimensionKeys"),
//   	},
//   	Entries: []interface{}{
//   		&LimitEntryProperty{
//   			Dimensions: map[string]*string{
//   				"dimensionsKey": jsii.String("dimensions"),
//   			},
//
//   			// the properties below are optional
//   			Connections: []interface{}{
//   				&RateConfigProperty{
//   					Period: jsii.String("period"),
//   					Rate: jsii.Number(123),
//   				},
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
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   	RateLimitId: jsii.String("rateLimitId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html
//
type CfnGatewayRateLimitProps struct {
	// Ordered list of dimension names defining the scope of this limit.
	//
	// Unique per gateway — no two limits can share the same dimensionKeys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html#cfn-bedrockagentcore-gatewayratelimit-dimensionkeys
	//
	DimensionKeys *[]*string `field:"required" json:"dimensionKeys" yaml:"dimensionKeys"`
	// Rule entries mapping dimension values to rate configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html#cfn-bedrockagentcore-gatewayratelimit-entries
	//
	Entries interface{} `field:"required" json:"entries" yaml:"entries"`
	// Optional human-readable description for this limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-gatewayratelimit.html#cfn-bedrockagentcore-gatewayratelimit-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
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

