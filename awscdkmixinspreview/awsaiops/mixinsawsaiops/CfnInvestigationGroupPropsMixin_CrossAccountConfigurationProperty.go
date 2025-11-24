package mixinsawsaiops


// This structure contains information about the cross-account configuration in the account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   crossAccountConfigurationProperty := &CrossAccountConfigurationProperty{
//   	SourceRoleArn: jsii.String("sourceRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-crossaccountconfiguration.html
//
type CfnInvestigationGroupPropsMixin_CrossAccountConfigurationProperty struct {
	// The ARN of an existing role which will be used to do investigations on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-crossaccountconfiguration.html#cfn-aiops-investigationgroup-crossaccountconfiguration-sourcerolearn
	//
	SourceRoleArn *string `field:"optional" json:"sourceRoleArn" yaml:"sourceRoleArn"`
}

