package awsrds


// Properties for defining a `CfnGlobalCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGlobalClusterProps := &cfnGlobalClusterProps{
//   	deletionProtection: jsii.Boolean(false),
//   	engine: jsii.String("engine"),
//   	engineVersion: jsii.String("engineVersion"),
//   	globalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	sourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	storageEncrypted: jsii.Boolean(false),
//   }
//
type CfnGlobalClusterProps struct {
	// The deletion protection setting for the new global database.
	//
	// The global database can't be deleted when deletion protection is enabled.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The name of the database engine to be used for this DB cluster.
	//
	// If this property isn't specified, the database engine is derived from the source DB cluster specified by the `SourceDBClusterIdentifier` property.
	//
	// > If the `SourceDBClusterIdentifier` property isn't specified, this property is required. If the `SourceDBClusterIdentifier` property is specified, make sure this property isn't specified.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The engine version of the Aurora global database.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The cluster identifier of the global database cluster.
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The DB cluster identifier or Amazon Resource Name (ARN) to use as the primary cluster of the global database.
	//
	// > If the `Engine` property isn't specified, this property is required. If the `Engine` property is specified, make sure this property isn't specified.
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// The storage encryption setting for the global database cluster.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
}

