package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDLPSetting`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDLPSettingProps := &CfnDLPSettingProps{
//   	DlpSettingId: jsii.String("dlpSettingId"),
//   	Enabled: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	ProviderConfig: &ProviderConfigProperty{
//   		MicrosoftPurview: &MicrosoftPurviewProviderConfigProperty{
//   			Credentials: &MicrosoftPurviewCredentialsProperty{
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   			LabelActionMappings: []interface{}{
//   				&LabelActionMappingProperty{
//   					Action: jsii.String("action"),
//   					LabelId: jsii.String("labelId"),
//   					LabelName: jsii.String("labelName"),
//   				},
//   			},
//   			UnmappedAction: jsii.String("unmappedAction"),
//   		},
//   	},
//   	ProviderOutageAction: jsii.String("providerOutageAction"),
//   	ProviderType: jsii.String("providerType"),
//
//   	// the properties below are optional
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html
//
type CfnDLPSettingProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-dlpsettingid
	//
	DlpSettingId *string `field:"required" json:"dlpSettingId" yaml:"dlpSettingId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-providerconfig
	//
	ProviderConfig interface{} `field:"required" json:"providerConfig" yaml:"providerConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-provideroutageaction
	//
	ProviderOutageAction *string `field:"required" json:"providerOutageAction" yaml:"providerOutageAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-providertype
	//
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

