package awscassandra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnKeyspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnKeyspaceProps := &CfnKeyspaceProps{
//   	ClientSideTimestampsEnabled: jsii.Boolean(false),
//   	KeyspaceName: jsii.String("keyspaceName"),
//   	ReplicationSpecification: &ReplicationSpecificationProperty{
//   		RegionList: []*string{
//   			jsii.String("regionList"),
//   		},
//   		ReplicationStrategy: jsii.String("replicationStrategy"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-keyspace.html
//
type CfnKeyspaceProps struct {
	// Indicates whether client-side timestamps are enabled (true) or disabled (false) for all tables in the keyspace.
	//
	// To add a Region to a single-Region keyspace with at least one table, the value must be set to true. After you've enabled client-side timestamps for a table, you can’t disable it again.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-keyspace.html#cfn-cassandra-keyspace-clientsidetimestampsenabled
	//
	ClientSideTimestampsEnabled interface{} `field:"optional" json:"clientSideTimestampsEnabled" yaml:"clientSideTimestampsEnabled"`
	// The name of the keyspace to be created.
	//
	// The keyspace name is case sensitive. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the keyspace name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// *Length constraints:* Minimum length of 1. Maximum length of 48.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-keyspace.html#cfn-cassandra-keyspace-keyspacename
	//
	KeyspaceName *string `field:"optional" json:"keyspaceName" yaml:"keyspaceName"`
	// Specifies the `ReplicationStrategy` of a keyspace. The options are:.
	//
	// - `SINGLE_REGION` for a single Region keyspace (optional) or
	// - `MULTI_REGION` for a multi-Region keyspace
	//
	// If no `ReplicationStrategy` is provided, the default is `SINGLE_REGION` . If you choose `MULTI_REGION` , you must also provide a `RegionList` with the AWS Regions that the keyspace is replicated in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-keyspace.html#cfn-cassandra-keyspace-replicationspecification
	//
	ReplicationSpecification interface{} `field:"optional" json:"replicationSpecification" yaml:"replicationSpecification"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cassandra-keyspace.html#cfn-cassandra-keyspace-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

