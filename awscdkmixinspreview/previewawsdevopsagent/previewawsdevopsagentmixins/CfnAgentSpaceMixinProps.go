package previewawsdevopsagentmixins


// Properties for CfnAgentSpacePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAgentSpaceMixinProps := &CfnAgentSpaceMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	OperatorApp: &OperatorAppProperty{
//   		Iam: &IamAuthConfigurationProperty{
//   			CreatedAt: jsii.String("createdAt"),
//   			OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//   			UpdatedAt: jsii.String("updatedAt"),
//   		},
//   		Idc: &IdcAuthConfigurationProperty{
//   			CreatedAt: jsii.String("createdAt"),
//   			IdcApplicationArn: jsii.String("idcApplicationArn"),
//   			IdcInstanceArn: jsii.String("idcInstanceArn"),
//   			OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//   			UpdatedAt: jsii.String("updatedAt"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html
//
type CfnAgentSpaceMixinProps struct {
	// The description of the Agent Space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the Agent Space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-agentspace.html#cfn-devopsagent-agentspace-operatorapp
	//
	OperatorApp interface{} `field:"optional" json:"operatorApp" yaml:"operatorApp"`
}

