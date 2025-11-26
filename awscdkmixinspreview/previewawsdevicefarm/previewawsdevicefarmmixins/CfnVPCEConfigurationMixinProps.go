package previewawsdevicefarmmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVPCEConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCEConfigurationMixinProps := &CfnVPCEConfigurationMixinProps{
//   	ServiceDnsName: jsii.String("serviceDnsName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpceConfigurationDescription: jsii.String("vpceConfigurationDescription"),
//   	VpceConfigurationName: jsii.String("vpceConfigurationName"),
//   	VpceServiceName: jsii.String("vpceServiceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-vpceconfiguration.html
//
type CfnVPCEConfigurationMixinProps struct {
	// The DNS name that Device Farm will use to map to the private service you want to access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-vpceconfiguration.html#cfn-devicefarm-vpceconfiguration-servicednsname
	//
	ServiceDnsName *string `field:"optional" json:"serviceDnsName" yaml:"serviceDnsName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-vpceconfiguration.html#cfn-devicefarm-vpceconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An optional description that provides details about your VPC endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-vpceconfiguration.html#cfn-devicefarm-vpceconfiguration-vpceconfigurationdescription
	//
	VpceConfigurationDescription *string `field:"optional" json:"vpceConfigurationDescription" yaml:"vpceConfigurationDescription"`
	// The friendly name you give to your VPC endpoint configuration to manage your configurations more easily.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-vpceconfiguration.html#cfn-devicefarm-vpceconfiguration-vpceconfigurationname
	//
	VpceConfigurationName *string `field:"optional" json:"vpceConfigurationName" yaml:"vpceConfigurationName"`
	// The name of the VPC endpoint service that you want to access from Device Farm.
	//
	// The name follows the format `com.amazonaws.vpce.us-west-2.vpce-svc-id` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devicefarm-vpceconfiguration.html#cfn-devicefarm-vpceconfiguration-vpceservicename
	//
	VpceServiceName *string `field:"optional" json:"vpceServiceName" yaml:"vpceServiceName"`
}

