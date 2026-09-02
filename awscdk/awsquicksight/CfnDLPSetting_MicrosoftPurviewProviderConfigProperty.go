package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   microsoftPurviewProviderConfigProperty := &MicrosoftPurviewProviderConfigProperty{
//   	Credentials: &MicrosoftPurviewCredentialsProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	LabelActionMappings: []interface{}{
//   		&LabelActionMappingProperty{
//   			Action: jsii.String("action"),
//   			LabelId: jsii.String("labelId"),
//   			LabelName: jsii.String("labelName"),
//   		},
//   	},
//   	UnmappedAction: jsii.String("unmappedAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-microsoftpurviewproviderconfig.html
//
type CfnDLPSetting_MicrosoftPurviewProviderConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-microsoftpurviewproviderconfig.html#cfn-quicksight-dlpsetting-microsoftpurviewproviderconfig-credentials
	//
	Credentials interface{} `field:"required" json:"credentials" yaml:"credentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-microsoftpurviewproviderconfig.html#cfn-quicksight-dlpsetting-microsoftpurviewproviderconfig-labelactionmappings
	//
	LabelActionMappings interface{} `field:"required" json:"labelActionMappings" yaml:"labelActionMappings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dlpsetting-microsoftpurviewproviderconfig.html#cfn-quicksight-dlpsetting-microsoftpurviewproviderconfig-unmappedaction
	//
	UnmappedAction *string `field:"required" json:"unmappedAction" yaml:"unmappedAction"`
}

