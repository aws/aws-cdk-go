package awseks


// The placement configuration for the etcd instances of your local Amazon EKS cluster on an AWS Outpost.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   etcdPlacementProperty := &EtcdPlacementProperty{
//   	SpreadLevel: jsii.String("spreadLevel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-etcdplacement.html
//
type CfnClusterPropsMixin_EtcdPlacementProperty struct {
	// Optional parameter to specify the placement group spread level for etcd instances.
	//
	// If not provided, EKS will deploy etcd instances without a placement group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-etcdplacement.html#cfn-eks-cluster-etcdplacement-spreadlevel
	//
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
}

