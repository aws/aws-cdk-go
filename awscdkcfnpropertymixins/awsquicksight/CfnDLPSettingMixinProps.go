package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDLPSettingPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDLPSettingMixinProps := &CfnDLPSettingMixinProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
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
type CfnDLPSettingMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-dlpsettingid
	//
	DlpSettingId *string `field:"optional" json:"dlpSettingId" yaml:"dlpSettingId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-enabled
	//
	// Default: - false.
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-providerconfig
	//
	ProviderConfig interface{} `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-provideroutageaction
	//
	ProviderOutageAction *string `field:"optional" json:"providerOutageAction" yaml:"providerOutageAction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-providertype
	//
	ProviderType *string `field:"optional" json:"providerType" yaml:"providerType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dlpsetting.html#cfn-quicksight-dlpsetting-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

