package awstimestream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInfluxDBInstance`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInfluxDBInstanceProps := &CfnInfluxDBInstanceProps{
//   	AllocatedStorage: jsii.Number(123),
//   	Bucket: jsii.String("bucket"),
//   	DbInstanceType: jsii.String("dbInstanceType"),
//   	DbParameterGroupIdentifier: jsii.String("dbParameterGroupIdentifier"),
//   	DbStorageType: jsii.String("dbStorageType"),
//   	DeploymentType: jsii.String("deploymentType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html
//
type CfnInfluxDBInstanceProps struct {
	// The amount of storage to allocate for your DB storage type in GiB (gibibytes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-allocatedstorage
	//
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// The name of the initial InfluxDB bucket.
	//
	// All InfluxDB data is stored in a bucket. A bucket combines the concept of a database and a retention period (the duration of time that each data point persists). A bucket belongs to an organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The Timestream for InfluxDB DB instance type to run on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-dbinstancetype
	//
	DbInstanceType *string `field:"optional" json:"dbInstanceType" yaml:"dbInstanceType"`
	// The name or id of the DB parameter group to assign to your DB instance.
	//
	// DB parameter groups specify how the database is configured. For example, DB parameter groups can specify the limit for query concurrency.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-dbparametergroupidentifier
	//
	DbParameterGroupIdentifier *string `field:"optional" json:"dbParameterGroupIdentifier" yaml:"dbParameterGroupIdentifier"`
	// The Timestream for InfluxDB DB storage type to read and write InfluxDB data.
	//
	// You can choose between 3 different types of provisioned Influx IOPS included storage according to your workloads requirements:
	//
	// - Influx IO Included 3000 IOPS
	// - Influx IO Included 12000 IOPS
	// - Influx IO Included 16000 IOPS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-dbstoragetype
	//
	DbStorageType *string `field:"optional" json:"dbStorageType" yaml:"dbStorageType"`
	// Specifies whether the Timestream for InfluxDB is deployed as Single-AZ or with a MultiAZ Standby for High availability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-deploymenttype
	//
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// Configuration for sending InfluxDB engine logs to a specified S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-logdeliveryconfiguration
	//
	LogDeliveryConfiguration interface{} `field:"optional" json:"logDeliveryConfiguration" yaml:"logDeliveryConfiguration"`
	// The name that uniquely identifies the DB instance when interacting with the Amazon Timestream for InfluxDB API and CLI commands.
	//
	// This name will also be a prefix included in the endpoint. DB instance names must be unique per customer and per region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Network type of the InfluxDB Instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-networktype
	//
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// The name of the initial organization for the initial admin user in InfluxDB.
	//
	// An InfluxDB organization is a workspace for a group of users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-organization
	//
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The password of the initial admin user created in InfluxDB.
	//
	// This password will allow you to access the InfluxDB UI to perform various administrative tasks and also use the InfluxDB CLI to create an operator token. These attributes will be stored in a Secret created in Amazon SecretManager in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The port number on which InfluxDB accepts connections.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Configures the DB instance with a public IP to facilitate access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-publiclyaccessible
	//
	// Default: - false.
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// A list of key-value pairs to associate with the DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The username of the initial admin user created in InfluxDB.
	//
	// Must start with a letter and can't end with a hyphen or contain two consecutive hyphens. For example, my-user1. This username will allow you to access the InfluxDB UI to perform various administrative tasks and also use the InfluxDB CLI to create an operator token. These attributes will be stored in a Secret created in Amazon Secrets Manager in your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
	// A list of VPC security group IDs to associate with the DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// A list of VPC subnet IDs to associate with the DB instance.
	//
	// Provide at least two VPC subnet IDs in different availability zones when deploying with a Multi-AZ standby.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-timestream-influxdbinstance.html#cfn-timestream-influxdbinstance-vpcsubnetids
	//
	VpcSubnetIds *[]*string `field:"optional" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
}

