package awsrds


// Properties for defining a `CfnGlobalCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGlobalClusterProps := &CfnGlobalClusterProps{
//   	DeletionProtection: jsii.Boolean(false),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	SourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	StorageEncrypted: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html
//
type CfnGlobalClusterProps struct {
	// Specifies whether to enable deletion protection for the new global database cluster.
	//
	// The global database can't be deleted when deletion protection is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The database engine to use for this global database cluster.
	//
	// Valid Values: `aurora-mysql | aurora-postgresql`
	//
	// Constraints:
	//
	// - Can't be specified if `SourceDBClusterIdentifier` is specified. In this case, Amazon Aurora uses the engine of the source DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The engine version to use for this global database cluster.
	//
	// Constraints:
	//
	// - Can't be specified if `SourceDBClusterIdentifier` is specified. In this case, Amazon Aurora uses the engine version of the source DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The cluster identifier for this global database cluster.
	//
	// This parameter is stored as a lowercase string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-globalclusteridentifier
	//
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The Amazon Resource Name (ARN) to use as the primary cluster of the global database.
	//
	// If you provide a value for this parameter, don't specify values for the following settings because Amazon Aurora uses the values from the specified source DB cluster:
	//
	// - `DatabaseName`
	// - `Engine`
	// - `EngineVersion`
	// - `StorageEncrypted`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-sourcedbclusteridentifier
	//
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// Specifies whether to enable storage encryption for the new global database cluster.
	//
	// Constraints:
	//
	// - Can't be specified if `SourceDBClusterIdentifier` is specified. In this case, Amazon Aurora uses the setting from the source DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html#cfn-rds-globalcluster-storageencrypted
	//
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
}

