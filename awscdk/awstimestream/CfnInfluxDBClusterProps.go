package awstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInfluxDBCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInfluxDBClusterProps := &CfnInfluxDBClusterProps{
//   	AllocatedStorage: jsii.Number(123),
//   	Bucket: jsii.String("bucket"),
//   	DbInstanceType: jsii.String("dbInstanceType"),
//   	DbParameterGroupIdentifier: jsii.String("dbParameterGroupIdentifier"),
//   	DbStorageType: jsii.String("dbStorageType"),
//   	DeploymentType: jsii.String("deploymentType"),
//   	FailoverMode: jsii.String("failoverMode"),
//   	LogDeliveryConfiguration: &LogDeliveryConfigurationProperty{
//   		S3Configuration: &S3ConfigurationProperty{
//   			BucketName: jsii.String("bucketName"),
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	NetworkType: jsii.String("networkType"),
//   	Organization: jsii.String("organization"),
//   	Password: jsii.String("password"),
//   	Port: jsii.Number(123),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Username: jsii.String("username"),
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   	VpcSubnetIds: []*string{
//   		jsii.String("vpcSubnetIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html
//
type CfnInfluxDBClusterProps struct {
	// The allocated storage for the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-allocatedstorage
	//
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// The bucket for the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The compute instance of the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-dbinstancetype
	//
	DbInstanceType *string `field:"optional" json:"dbInstanceType" yaml:"dbInstanceType"`
	// The name of an existing InfluxDB parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-dbparametergroupidentifier
	//
	DbParameterGroupIdentifier *string `field:"optional" json:"dbParameterGroupIdentifier" yaml:"dbParameterGroupIdentifier"`
	// The storage type of the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-dbstoragetype
	//
	DbStorageType *string `field:"optional" json:"dbStorageType" yaml:"dbStorageType"`
	// Deployment type of the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-deploymenttype
	//
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// Failover mode of the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-failovermode
	//
	FailoverMode *string `field:"optional" json:"failoverMode" yaml:"failoverMode"`
	// Configuration for sending logs to customer account from the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-logdeliveryconfiguration
	//
	LogDeliveryConfiguration interface{} `field:"optional" json:"logDeliveryConfiguration" yaml:"logDeliveryConfiguration"`
	// The unique name that is associated with the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Network type of the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The organization for the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-organization
	//
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The password for the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port number on which InfluxDB accepts connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Attach a public IP to the customer ENI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-publiclyaccessible
	//
	// Default: - false.
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An arbitrary set of tags (key-value pairs) for this DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The username for the InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
	// A list of Amazon EC2 VPC security groups to associate with this InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// A list of EC2 subnet IDs for this InfluxDB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbcluster.html#cfn-timestream-influxdbcluster-vpcsubnetids
	//
	VpcSubnetIds *[]*string `field:"optional" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
}

