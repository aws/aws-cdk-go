package awscodeconnections

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnHostPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnHostMixinProps := &CfnHostMixinProps{
//   	Name: jsii.String("name"),
//   	ProviderEndpoint: jsii.String("providerEndpoint"),
//   	ProviderType: jsii.String("providerType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConfiguration: &VpcConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		TlsCertificate: jsii.String("tlsCertificate"),
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-host.html
//
type CfnHostMixinProps struct {
	// The name of the host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-host.html#cfn-codeconnections-host-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The endpoint of the infrastructure where your provider type is installed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-host.html#cfn-codeconnections-host-providerendpoint
	//
	ProviderEndpoint *string `field:"optional" json:"providerEndpoint" yaml:"providerEndpoint"`
	// The name of the installed provider to be associated with your connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-host.html#cfn-codeconnections-host-providertype
	//
	ProviderType *string `field:"optional" json:"providerType" yaml:"providerType"`
	// Tags to apply to the host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-host.html#cfn-codeconnections-host-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The VPC configuration provisioned for the host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeconnections-host.html#cfn-codeconnections-host-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

