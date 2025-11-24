package mixinsawssecuritylake


// Provides replication configuration details for objects stored in the Amazon Security Lake data lake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   replicationConfigurationProperty := &ReplicationConfigurationProperty{
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-replicationconfiguration.html
//
type CfnDataLakePropsMixin_ReplicationConfigurationProperty struct {
	// Specifies one or more centralized rollup Regions.
	//
	// The AWS Region specified in the region parameter of the `CreateDataLake` or `UpdateDataLake` operations contributes data to the rollup Region or Regions specified in this parameter.
	//
	// Replication enables automatic, asynchronous copying of objects across Amazon S3 buckets. S3 buckets that are configured for object replication can be owned by the same AWS account or by different accounts. You can replicate objects to a single destination bucket or to multiple destination buckets. The destination buckets can be in different Regions or within the same Region as the source bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-replicationconfiguration.html#cfn-securitylake-datalake-replicationconfiguration-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Replication settings for the Amazon S3 buckets.
	//
	// This parameter uses the AWS Identity and Access Management (IAM) role you created that is managed by Security Lake , to ensure the replication setting is correct.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-replicationconfiguration.html#cfn-securitylake-datalake-replicationconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

