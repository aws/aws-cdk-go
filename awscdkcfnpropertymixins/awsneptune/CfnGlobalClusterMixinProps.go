package awsneptune

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnGlobalClusterPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnGlobalClusterMixinProps := &CfnGlobalClusterMixinProps{
//   	DeletionProtection: jsii.Boolean(false),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	SourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	StorageEncrypted: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-globalcluster.html
//
type CfnGlobalClusterMixinProps struct {
	// Whether deletion protection is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-globalcluster.html#cfn-neptune-globalcluster-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The name of the database engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-globalcluster.html#cfn-neptune-globalcluster-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The version number of the database engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-globalcluster.html#cfn-neptune-globalcluster-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The cluster identifier of the global database cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-globalcluster.html#cfn-neptune-globalcluster-globalclusteridentifier
	//
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The Amazon Resource Name (ARN) of an existing Neptune DB cluster to use as the primary cluster of the new global database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-globalcluster.html#cfn-neptune-globalcluster-sourcedbclusteridentifier
	//
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// Whether the global database cluster is storage encrypted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-globalcluster.html#cfn-neptune-globalcluster-storageencrypted
	//
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-globalcluster.html#cfn-neptune-globalcluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

