package awsdevopsagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   operatorAppProperty := &OperatorAppProperty{
//   	Iam: &IamAuthConfigurationProperty{
//   		OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//
//   		// the properties below are optional
//   		CreatedAt: jsii.String("createdAt"),
//   		UpdatedAt: jsii.String("updatedAt"),
//   	},
//   	Idc: &IdcAuthConfigurationProperty{
//   		IdcInstanceArn: jsii.String("idcInstanceArn"),
//   		OperatorAppRoleArn: jsii.String("operatorAppRoleArn"),
//
//   		// the properties below are optional
//   		CreatedAt: jsii.String("createdAt"),
//   		IdcApplicationArn: jsii.String("idcApplicationArn"),
//   		UpdatedAt: jsii.String("updatedAt"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-operatorapp.html
//
type CfnAgentSpace_OperatorAppProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-operatorapp.html#cfn-devopsagent-agentspace-operatorapp-iam
	//
	Iam interface{} `field:"optional" json:"iam" yaml:"iam"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-agentspace-operatorapp.html#cfn-devopsagent-agentspace-operatorapp-idc
	//
	Idc interface{} `field:"optional" json:"idc" yaml:"idc"`
}

