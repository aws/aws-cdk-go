package mixinsawselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnParameterGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnParameterGroupMixinProps := &CfnParameterGroupMixinProps{
//   	CacheParameterGroupFamily: jsii.String("cacheParameterGroupFamily"),
//   	Description: jsii.String("description"),
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html
//
type CfnParameterGroupMixinProps struct {
	// The name of the cache parameter group family that this cache parameter group is compatible with.
	//
	// Valid values are: `valkey8` | `valkey7` | `memcached1.4` | `memcached1.5` | `memcached1.6` | `redis2.6` | `redis2.8` | `redis3.2` | `redis4.0` | `redis5.0` | `redis6.x` | `redis7`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html#cfn-elasticache-parametergroup-cacheparametergroupfamily
	//
	CacheParameterGroupFamily *string `field:"optional" json:"cacheParameterGroupFamily" yaml:"cacheParameterGroupFamily"`
	// The description for this cache parameter group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html#cfn-elasticache-parametergroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A comma-delimited list of parameter name/value pairs.
	//
	// For example:
	//
	// ```
	// "Properties" : { "cas_disabled" : "1", "chunk_size_growth_factor" : "1.02"
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html#cfn-elasticache-parametergroup-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// A tag that can be added to an ElastiCache parameter group.
	//
	// Tags are composed of a Key/Value pair. You can use tags to categorize and track all your parameter groups. A tag with a null Value is permitted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-parametergroup.html#cfn-elasticache-parametergroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

