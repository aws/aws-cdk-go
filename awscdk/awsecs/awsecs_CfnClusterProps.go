package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &cfnClusterProps{
//   	capacityProviders: []*string{
//   		jsii.String("capacityProviders"),
//   	},
//   	clusterName: jsii.String("clusterName"),
//   	clusterSettings: []interface{}{
//   		&clusterSettingsProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	configuration: &clusterConfigurationProperty{
//   		executeCommandConfiguration: &executeCommandConfigurationProperty{
//   			kmsKeyId: jsii.String("kmsKeyId"),
//   			logConfiguration: &executeCommandLogConfigurationProperty{
//   				cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   				cloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   				s3BucketName: jsii.String("s3BucketName"),
//   				s3EncryptionEnabled: jsii.Boolean(false),
//   				s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   			},
//   			logging: jsii.String("logging"),
//   		},
//   	},
//   	defaultCapacityProviderStrategy: []interface{}{
//   		&capacityProviderStrategyItemProperty{
//   			base: jsii.Number(123),
//   			capacityProvider: jsii.String("capacityProvider"),
//   			weight: jsii.Number(123),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnClusterProps struct {
	// The short name of one or more capacity providers to associate with the cluster.
	//
	// A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy.
	//
	// If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must already be created and not already associated with another cluster.
	//
	// To use an AWS Fargate capacity provider, specify either the `FARGATE` or `FARGATE_SPOT` capacity providers. The AWS Fargate capacity providers are available to all accounts and only need to be associated with a cluster to be used.
	CapacityProviders *[]*string `field:"optional" json:"capacityProviders" yaml:"capacityProviders"`
	// A user-generated string that you use to identify your cluster.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the name.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// The setting to use when creating a cluster.
	//
	// This parameter is used to enable CloudWatch Container Insights for a cluster. If this value is specified, it will override the `containerInsights` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html) .
	ClusterSettings interface{} `field:"optional" json:"clusterSettings" yaml:"clusterSettings"`
	// The execute command configuration for the cluster.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The default capacity provider strategy for the cluster.
	//
	// When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.
	DefaultCapacityProviderStrategy interface{} `field:"optional" json:"defaultCapacityProviderStrategy" yaml:"defaultCapacityProviderStrategy"`
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
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	// - Tag keys and values are case-sensitive.
	// - Do not use `aws:` , `AWS:` , or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

