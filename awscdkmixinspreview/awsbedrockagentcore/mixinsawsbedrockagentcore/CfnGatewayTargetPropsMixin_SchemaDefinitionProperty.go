package mixinsawsbedrockagentcore


// The schema definition for the gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var schemaDefinitionProperty_ SchemaDefinitionProperty
//
//   schemaDefinitionProperty := &SchemaDefinitionProperty{
//   	Description: jsii.String("description"),
//   	Items: schemaDefinitionProperty_,
//   	Properties: map[string]interface{}{
//   		"propertiesKey": schemaDefinitionProperty_,
//   	},
//   	Required: []*string{
//   		jsii.String("required"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.html
//
type CfnGatewayTargetPropsMixin_SchemaDefinitionProperty struct {
	// The workload identity details for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.html#cfn-bedrockagentcore-gatewaytarget-schemadefinition-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.html#cfn-bedrockagentcore-gatewaytarget-schemadefinition-items
	//
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// The schema definition properties for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.html#cfn-bedrockagentcore-gatewaytarget-schemadefinition-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The schema definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.html#cfn-bedrockagentcore-gatewaytarget-schemadefinition-required
	//
	Required *[]*string `field:"optional" json:"required" yaml:"required"`
	// The scheme definition type for the gateway target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-schemadefinition.html#cfn-bedrockagentcore-gatewaytarget-schemadefinition-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

