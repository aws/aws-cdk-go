package mixinsawsbedrockagentcore


// The tool schema for the gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   toolSchemaProperty := &ToolSchemaProperty{
//   	InlinePayload: []interface{}{
//   		&ToolDefinitionProperty{
//   			Description: jsii.String("description"),
//   			InputSchema: &SchemaDefinitionProperty{
//   				Description: jsii.String("description"),
//   				Items: schemaDefinitionProperty_,
//   				Properties: map[string]interface{}{
//   					"propertiesKey": schemaDefinitionProperty_,
//   				},
//   				Required: []*string{
//   					jsii.String("required"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			Name: jsii.String("name"),
//   			OutputSchema: &SchemaDefinitionProperty{
//   				Description: jsii.String("description"),
//   				Items: schemaDefinitionProperty_,
//   				Properties: map[string]interface{}{
//   					"propertiesKey": schemaDefinitionProperty_,
//   				},
//   				Required: []*string{
//   					jsii.String("required"),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	S3: &S3ConfigurationProperty{
//   		BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   		Uri: jsii.String("uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-toolschema.html
//
type CfnGatewayTargetPropsMixin_ToolSchemaProperty struct {
	// The inline payload for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-toolschema.html#cfn-bedrockagentcore-gatewaytarget-toolschema-inlinepayload
	//
	InlinePayload interface{} `field:"optional" json:"inlinePayload" yaml:"inlinePayload"`
	// The S3 tool schema for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-toolschema.html#cfn-bedrockagentcore-gatewaytarget-toolschema-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

