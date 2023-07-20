package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &CfnClusterProps{
//   	CapacityProviders: []*string{
//   		jsii.String("capacityProviders"),
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	ClusterSettings: []interface{}{
//   		&ClusterSettingsProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Configuration: &ClusterConfigurationProperty{
//   		ExecuteCommandConfiguration: &ExecuteCommandConfigurationProperty{
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			LogConfiguration: &ExecuteCommandLogConfigurationProperty{
//   				CloudWatchEncryptionEnabled: jsii.Boolean(false),
//   				CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				S3BucketName: jsii.String("s3BucketName"),
//   				S3EncryptionEnabled: jsii.Boolean(false),
//   				S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   			},
//   			Logging: jsii.String("logging"),
//   		},
//   	},
//   	DefaultCapacityProviderStrategy: []interface{}{
//   		&CapacityProviderStrategyItemProperty{
//   			Base: jsii.Number(123),
//   			CapacityProvider: jsii.String("capacityProvider"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	ServiceConnectDefaults: &ServiceConnectDefaultsProperty{
//   		Namespace: jsii.String("namespace"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html
//
type CfnClusterProps struct {
	// The short name of one or more capacity providers to associate with the cluster.
	//
	// A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy when calling the [CreateService](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html) or [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) actions.
	//
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must be created but not associated with another cluster. New Auto Scaling group capacity providers can be created with the [CreateCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCapacityProvider.html) API operation.
	//
	// To use a AWS Fargate capacity provider, specify either the `FARGATE` or `FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers are available to all accounts and only need to be associated with a cluster to be used.
	//
	// The [PutCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutCapacityProvider.html) API operation is used to update the list of available capacity providers for a cluster after the cluster is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-capacityproviders
	//
	CapacityProviders *[]*string `field:"optional" json:"capacityProviders" yaml:"capacityProviders"`
	// A user-generated string that you use to identify your cluster.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The settings to use when creating a cluster.
	//
	// This parameter is used to turn on CloudWatch Container Insights for a cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-clustersettings
	//
	ClusterSettings interface{} `field:"optional" json:"clusterSettings" yaml:"clusterSettings"`
	// The execute command configuration for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The default capacity provider strategy for the cluster.
	//
	// When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-defaultcapacityproviderstrategy
	//
	DefaultCapacityProviderStrategy interface{} `field:"optional" json:"defaultCapacityProviderStrategy" yaml:"defaultCapacityProviderStrategy"`
	// Use this parameter to set a default Service Connect namespace.
	//
	// After you set a default Service Connect namespace, any new services with Service Connect turned on that are created in the cluster are added as client services in the namespace. This setting only applies to new services that set the `enabled` parameter to `true` in the `ServiceConnectConfiguration` . You can set the namespace of each service individually in the `ServiceConnectConfiguration` to override this default parameter.
	//
	// Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-serviceconnectdefaults
	//
	ServiceConnectDefaults interface{} `field:"optional" json:"serviceConnectDefaults" yaml:"serviceConnectDefaults"`
	// The metadata that you apply to the cluster to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value. You define both.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource - 50
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length - 128 Unicode characters in UTF-8
	// - Maximum value length - 256 Unicode characters in UTF-8
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-cluster.html#cfn-ecs-cluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

