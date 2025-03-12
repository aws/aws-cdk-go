package awsemrcontainers


// The information about the container provider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerProviderProperty := &ContainerProviderProperty{
//   	Id: jsii.String("id"),
//   	Info: &ContainerInfoProperty{
//   		EksInfo: &EksInfoProperty{
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-containerprovider.html
//
type CfnVirtualCluster_ContainerProviderProperty struct {
	// The ID of the container cluster.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 100
	//
	// *Pattern* : `^[0-9A-Za-z][A-Za-z0-9\-_]*`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-containerprovider.html#cfn-emrcontainers-virtualcluster-containerprovider-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The information about the container cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-containerprovider.html#cfn-emrcontainers-virtualcluster-containerprovider-info
	//
	Info interface{} `field:"required" json:"info" yaml:"info"`
	// The type of the container provider.
	//
	// Amazon EKS is the only supported type as of now.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-containerprovider.html#cfn-emrcontainers-virtualcluster-containerprovider-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

