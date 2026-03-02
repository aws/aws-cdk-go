package previewawsdevopsagentmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iamAuthConfigurationProperty := &IamAuthConfigurationProperty{
//   	CreatedAt: jsii.String("createdAt"),
//   	OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-iamauthconfiguration.html
//
type CfnAgentSpacePropsMixin_IamAuthConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-iamauthconfiguration.html#cfn-devopsagent-agentspace-iamauthconfiguration-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-iamauthconfiguration.html#cfn-devopsagent-agentspace-iamauthconfiguration-operatorapprolearn
	//
	OperatorAppRoleArn *string `field:"optional" json:"operatorAppRoleArn" yaml:"operatorAppRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-iamauthconfiguration.html#cfn-devopsagent-agentspace-iamauthconfiguration-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

