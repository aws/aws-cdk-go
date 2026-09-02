package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   providerConfigProperty := &ProviderConfigProperty{
//   	MicrosoftPurview: &MicrosoftPurviewProviderConfigProperty{
//   		Credentials: &MicrosoftPurviewCredentialsProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		LabelActionMappings: []interface{}{
//   			&LabelActionMappingProperty{
//   				Action: jsii.String("action"),
//   				LabelId: jsii.String("labelId"),
//   				LabelName: jsii.String("labelName"),
//   			},
//   		},
//   		UnmappedAction: jsii.String("unmappedAction"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-providerconfig.html
//
type CfnDLPSetting_ProviderConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-providerconfig.html#cfn-quicksight-dlpsetting-providerconfig-microsoftpurview
	//
	MicrosoftPurview interface{} `field:"required" json:"microsoftPurview" yaml:"microsoftPurview"`
}

