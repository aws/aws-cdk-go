package awsmsk


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   zookeeperAccessProperty := &ZookeeperAccessProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-zookeeperaccess.html
//
type CfnClusterPropsMixin_ZookeeperAccessProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-zookeeperaccess.html#cfn-msk-cluster-zookeeperaccess-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

