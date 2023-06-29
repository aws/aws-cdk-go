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
type CfnKeyspaceProps struct {
	// The name of the keyspace to be created.
	//
	// The keyspace name is case sensitive. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the keyspace name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// *Length constraints:* Minimum length of 3. Maximum length of 255.
	//
	// *Pattern:* `^[a-zA-Z0-9][a-zA-Z0-9_]{1,47}$`.
	KeyspaceName *string `field:"optional" json:"keyspaceName" yaml:"keyspaceName"`
	// Specifies the `ReplicationStrategy` of a keyspace. The options are:.
	//
	// - `SINGLE_REGION` for a single Region keyspace (optional) or
	// - `MULTI_REGION` for a multi-Region keyspace
	//
	// If no `ReplicationStrategy` is provided, the default is `SINGLE_REGION` . If you choose `MULTI_REGION` , you must also provide a `RegionList` with the AWS Regions that the keyspace is replicated in.
	ReplicationSpecification interface{} `field:"optional" json:"replicationSpecification" yaml:"replicationSpecification"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

