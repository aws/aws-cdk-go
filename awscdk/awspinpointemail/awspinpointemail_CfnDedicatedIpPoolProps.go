package awspinpointemail


// Properties for defining a `CfnDedicatedIpPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDedicatedIpPoolProps := &cfnDedicatedIpPoolProps{
//   	poolName: jsii.String("poolName"),
//   	tags: []tagsProperty{
//   		&tagsProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDedicatedIpPoolProps struct {
	// The name of the dedicated IP pool.
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// An object that defines the tags (keys and values) that you want to associate with the dedicated IP pool.
	Tags *[]*CfnDedicatedIpPool_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

