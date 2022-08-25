package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
)

// The properties to import from the ECS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var autoScalingGroup autoScalingGroup
//   var bucket bucket
//   var key key
//   var logGroup logGroup
//   var namespace iNamespace
//   var securityGroup securityGroup
//   var vpc vpc
//
//   clusterAttributes := &clusterAttributes{
//   	clusterName: jsii.String("clusterName"),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	vpc: vpc,
//
//   	// the properties below are optional
//   	autoscalingGroup: autoScalingGroup,
//   	clusterArn: jsii.String("clusterArn"),
//   	defaultCloudMapNamespace: namespace,
//   	executeCommandConfiguration: &executeCommandConfiguration{
//   		kmsKey: key,
//   		logConfiguration: &executeCommandLogConfiguration{
//   			cloudWatchEncryptionEnabled: jsii.Boolean(false),
//   			cloudWatchLogGroup: logGroup,
//   			s3Bucket: bucket,
//   			s3EncryptionEnabled: jsii.Boolean(false),
//   			s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//   		logging: awscdk.Aws_ecs.executeCommandLogging_NONE,
//   	},
//   	hasEc2Capacity: jsii.Boolean(false),
//   }
//
type ClusterAttributes struct {
	// The name of the cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The security groups associated with the container instances registered to the cluster.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The VPC associated with the cluster.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Autoscaling group added to the cluster if capacity is added.
	AutoscalingGroup awsautoscaling.IAutoScalingGroup `field:"optional" json:"autoscalingGroup" yaml:"autoscalingGroup"`
	// The Amazon Resource Name (ARN) that identifies the cluster.
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// The AWS Cloud Map namespace to associate with the cluster.
	DefaultCloudMapNamespace awsservicediscovery.INamespace `field:"optional" json:"defaultCloudMapNamespace" yaml:"defaultCloudMapNamespace"`
	// The execute command configuration for the cluster.
	ExecuteCommandConfiguration *ExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// Specifies whether the cluster has EC2 instance capacity.
	HasEc2Capacity *bool `field:"optional" json:"hasEc2Capacity" yaml:"hasEc2Capacity"`
}

