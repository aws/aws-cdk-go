package previewawsemrcontainersmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnEndpointPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var eMREKSConfigurationProperty_ EMREKSConfigurationProperty
//
//   cfnEndpointMixinProps := &CfnEndpointMixinProps{
//   	ConfigurationOverrides: &ConfigurationOverridesProperty{
//   		ApplicationConfiguration: []interface{}{
//   			&EMREKSConfigurationProperty{
//   				Classification: jsii.String("classification"),
//   				Configurations: []interface{}{
//   					eMREKSConfigurationProperty_,
//   				},
//   				Properties: map[string]*string{
//   					"propertiesKey": jsii.String("properties"),
//   				},
//   			},
//   		},
//   		MonitoringConfiguration: &MonitoringConfigurationProperty{
//   			CloudWatchMonitoringConfiguration: &CloudWatchMonitoringConfigurationProperty{
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   			},
//   			ContainerLogRotationConfiguration: &ContainerLogRotationConfigurationProperty{
//   				MaxFilesToKeep: jsii.Number(123),
//   				RotationSize: jsii.String("rotationSize"),
//   			},
//   			PersistentAppUi: jsii.String("persistentAppUi"),
//   			S3MonitoringConfiguration: &S3MonitoringConfigurationProperty{
//   				LogUri: jsii.String("logUri"),
//   			},
//   		},
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Name: jsii.String("name"),
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   	VirtualClusterId: jsii.String("virtualClusterId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html
//
type CfnEndpointMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-configurationoverrides
	//
	ConfigurationOverrides interface{} `field:"optional" json:"configurationOverrides" yaml:"configurationOverrides"`
	// The execution role ARN for the managed endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the managed endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon EMR release label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-releaselabel
	//
	ReleaseLabel *string `field:"optional" json:"releaseLabel" yaml:"releaseLabel"`
	// An array of key-value pairs to apply to this managed endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of the managed endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The ID of the virtual cluster for which the managed endpoint is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-virtualclusterid
	//
	VirtualClusterId *string `field:"optional" json:"virtualClusterId" yaml:"virtualClusterId"`
}

