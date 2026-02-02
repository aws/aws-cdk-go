package awsemrcontainers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eMREKSConfigurationProperty_ EMREKSConfigurationProperty
//
//   cfnEndpointProps := &CfnEndpointProps{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	ReleaseLabel: jsii.String("releaseLabel"),
//   	Type: jsii.String("type"),
//   	VirtualClusterId: jsii.String("virtualClusterId"),
//
//   	// the properties below are optional
//   	ConfigurationOverrides: &ConfigurationOverridesProperty{
//   		ApplicationConfiguration: []interface{}{
//   			&EMREKSConfigurationProperty{
//   				Classification: jsii.String("classification"),
//
//   				// the properties below are optional
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
//
//   				// the properties below are optional
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
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html
//
type CfnEndpointProps struct {
	// The execution role ARN for the managed endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-executionrolearn
	//
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The Amazon EMR release label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-releaselabel
	//
	ReleaseLabel *string `field:"required" json:"releaseLabel" yaml:"releaseLabel"`
	// The type of the managed endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The ID of the virtual cluster for which the managed endpoint is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-virtualclusterid
	//
	VirtualClusterId *string `field:"required" json:"virtualClusterId" yaml:"virtualClusterId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-configurationoverrides
	//
	ConfigurationOverrides interface{} `field:"optional" json:"configurationOverrides" yaml:"configurationOverrides"`
	// The name of the managed endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this managed endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emrcontainers-endpoint.html#cfn-emrcontainers-endpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

