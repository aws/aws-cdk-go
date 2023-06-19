package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnParameterGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnParameterGroupProps := &CfnParameterGroupProps{
//   	CacheParameterGroupFamily: jsii.String("cacheParameterGroupFamily"),
//   	Description: jsii.String("description"),
//
//   	// the properties below are optional
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnParameterGroupProps struct {
	// The name of the cache parameter group family that this cache parameter group is compatible with.
	//
	// Valid values are: `memcached1.4` | `memcached1.5` | `memcached1.6` | `redis2.6` | `redis2.8` | `redis3.2` | `redis4.0` | `redis5.0` | `redis6.x` | `redis7`
	CacheParameterGroupFamily *string `field:"required" json:"cacheParameterGroupFamily" yaml:"cacheParameterGroupFamily"`
	// The description for this cache parameter group.
	Description *string `field:"required" json:"description" yaml:"description"`
	// A comma-delimited list of parameter name/value pairs.
	//
	// For example:
	//
	// ```
	// "Properties" : { "cas_disabled" : "1", "chunk_size_growth_factor" : "1.02"
	// }
	// ```.
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// A tag that can be added to an ElastiCache parameter group.
	//
	// Tags are composed of a Key/Value pair. You can use tags to categorize and track all your parameter groups. A tag with a null Value is permitted.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

