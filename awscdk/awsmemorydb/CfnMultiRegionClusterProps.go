package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMultiRegionCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMultiRegionClusterProps := &CfnMultiRegionClusterProps{
//   	NodeType: jsii.String("nodeType"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	MultiRegionClusterNameSuffix: jsii.String("multiRegionClusterNameSuffix"),
//   	MultiRegionParameterGroupName: jsii.String("multiRegionParameterGroupName"),
//   	NumShards: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TlsEnabled: jsii.Boolean(false),
//   	UpdateStrategy: jsii.String("updateStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html
//
type CfnMultiRegionClusterProps struct {
	// The node type used by the multi-Region cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-nodetype
	//
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// The description of the multi-Region cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the engine used by the multi-Region cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The version of the engine used by the multi-Region cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The name of the Multi Region cluster.
	//
	// This value must be unique as it also serves as the multi region cluster identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-multiregionclusternamesuffix
	//
	MultiRegionClusterNameSuffix *string `field:"optional" json:"multiRegionClusterNameSuffix" yaml:"multiRegionClusterNameSuffix"`
	// The name of the multi-Region parameter group associated with the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-multiregionparametergroupname
	//
	MultiRegionParameterGroupName *string `field:"optional" json:"multiRegionParameterGroupName" yaml:"multiRegionParameterGroupName"`
	// The number of shards the multi region cluster will contain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-numshards
	//
	NumShards *float64 `field:"optional" json:"numShards" yaml:"numShards"`
	// An array of key-value pairs to apply to this multi region cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indiciates if the multi-Region cluster is TLS enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-tlsenabled
	//
	TlsEnabled interface{} `field:"optional" json:"tlsEnabled" yaml:"tlsEnabled"`
	// An enum string value that determines the update strategy for scaling.
	//
	// Possible values are 'COORDINATED' and 'UNCOORDINATED'. Default is 'COORDINATED'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html#cfn-memorydb-multiregioncluster-updatestrategy
	//
	UpdateStrategy *string `field:"optional" json:"updateStrategy" yaml:"updateStrategy"`
}

