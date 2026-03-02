package awsdevopsagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iamAuthConfigurationProperty := &IamAuthConfigurationProperty{
//   	OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-iamauthconfiguration.html
//
type CfnAgentSpace_IamAuthConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-iamauthconfiguration.html#cfn-devopsagent-agentspace-iamauthconfiguration-operatorapprolearn
	//
	OperatorAppRoleArn *string `field:"required" json:"operatorAppRoleArn" yaml:"operatorAppRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-iamauthconfiguration.html#cfn-devopsagent-agentspace-iamauthconfiguration-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-iamauthconfiguration.html#cfn-devopsagent-agentspace-iamauthconfiguration-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}

