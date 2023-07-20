package awspinpointemail


// Properties for defining a `CfnDedicatedIpPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDedicatedIpPoolProps := &CfnDedicatedIpPoolProps{
//   	PoolName: jsii.String("poolName"),
//   	Tags: []tagsProperty{
//   		&tagsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-dedicatedippool.html
//
type CfnDedicatedIpPoolProps struct {
	// The name of the dedicated IP pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-dedicatedippool.html#cfn-pinpointemail-dedicatedippool-poolname
	//
	PoolName *string `field:"optional" json:"poolName" yaml:"poolName"`
	// An object that defines the tags (keys and values) that you want to associate with the dedicated IP pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-dedicatedippool.html#cfn-pinpointemail-dedicatedippool-tags
	//
	Tags *[]*CfnDedicatedIpPool_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

