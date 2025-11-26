package previewawsemrcontainersmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVirtualClusterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVirtualClusterMixinProps := &CfnVirtualClusterMixinProps{
//   	ContainerProvider: &ContainerProviderProperty{
//   		Id: jsii.String("id"),
//   		Info: &ContainerInfoProperty{
//   			EksInfo: &EksInfoProperty{
//   				Namespace: jsii.String("namespace"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Name: jsii.String("name"),
//   	SecurityConfigurationId: jsii.String("securityConfigurationId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-virtualcluster.html
//
type CfnVirtualClusterMixinProps struct {
	// The container provider of the virtual cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-virtualcluster.html#cfn-emrcontainers-virtualcluster-containerprovider
	//
	ContainerProvider interface{} `field:"optional" json:"containerProvider" yaml:"containerProvider"`
	// The name of the virtual cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-virtualcluster.html#cfn-emrcontainers-virtualcluster-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ID of the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-virtualcluster.html#cfn-emrcontainers-virtualcluster-securityconfigurationid
	//
	SecurityConfigurationId *string `field:"optional" json:"securityConfigurationId" yaml:"securityConfigurationId"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-virtualcluster.html#cfn-emrcontainers-virtualcluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

