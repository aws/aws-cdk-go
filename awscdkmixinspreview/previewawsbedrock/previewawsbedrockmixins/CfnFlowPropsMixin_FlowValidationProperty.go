package previewawsbedrockmixins


// Contains information about validation of the flow.
//
// This data type is used in the following API operations:
//
// - [GetFlow response](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetFlow.html#API_agent_GetFlow_ResponseSyntax)
// - [GetFlowVersion response](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetFlowVersion.html#API_agent_GetFlowVersion_ResponseSyntax)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   flowValidationProperty := &FlowValidationProperty{
//   	Message: jsii.String("message"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowvalidation.html
//
type CfnFlowPropsMixin_FlowValidationProperty struct {
	// A message describing the validation error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowvalidation.html#cfn-bedrock-flow-flowvalidation-message
	//
	Message *string `field:"optional" json:"message" yaml:"message"`
}

