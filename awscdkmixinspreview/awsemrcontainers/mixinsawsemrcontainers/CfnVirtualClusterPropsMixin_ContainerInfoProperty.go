package mixinsawsemrcontainers


// The information about the container used for a job run or a managed endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerInfoProperty := &ContainerInfoProperty{
//   	EksInfo: &EksInfoProperty{
//   		Namespace: jsii.String("namespace"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-containerinfo.html
//
type CfnVirtualClusterPropsMixin_ContainerInfoProperty struct {
	// The information about the Amazon EKS cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-containerinfo.html#cfn-emrcontainers-virtualcluster-containerinfo-eksinfo
	//
	EksInfo interface{} `field:"optional" json:"eksInfo" yaml:"eksInfo"`
}

