package previewawsdevopsagentmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idcAuthConfigurationProperty := &IdcAuthConfigurationProperty{
//   	CreatedAt: jsii.String("createdAt"),
//   	IdcApplicationArn: jsii.String("idcApplicationArn"),
//   	IdcInstanceArn: jsii.String("idcInstanceArn"),
//   	OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idcauthconfiguration.html
//
type CfnAgentSpacePropsMixin_IdcAuthConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idcauthconfiguration.html#cfn-devopsagent-agentspace-idcauthconfiguration-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idcauthconfiguration.html#cfn-devopsagent-agentspace-idcauthconfiguration-idcapplicationarn
	//
	IdcApplicationArn *string `field:"optional" json:"idcApplicationArn" yaml:"idcApplicationArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idcauthconfiguration.html#cfn-devopsagent-agentspace-idcauthconfiguration-idcinstancearn
	//
	IdcInstanceArn *string `field:"optional" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idcauthconfiguration.html#cfn-devopsagent-agentspace-idcauthconfiguration-operatorapprolearn
	//
	OperatorAppRoleArn *string `field:"optional" json:"operatorAppRoleArn" yaml:"operatorAppRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-idcauthconfiguration.html#cfn-devopsagent-agentspace-idcauthconfiguration-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

