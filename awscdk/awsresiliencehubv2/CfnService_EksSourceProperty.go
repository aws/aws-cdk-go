package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eksSourceProperty := &EksSourceProperty{
//   	ClusterArn: jsii.String("clusterArn"),
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekssource.html
//
type CfnService_EksSourceProperty struct {
	// ARN of the EKS cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekssource.html#cfn-resiliencehubv2-service-ekssource-clusterarn
	//
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// EKS namespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-ekssource.html#cfn-resiliencehubv2-service-ekssource-namespaces
	//
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

