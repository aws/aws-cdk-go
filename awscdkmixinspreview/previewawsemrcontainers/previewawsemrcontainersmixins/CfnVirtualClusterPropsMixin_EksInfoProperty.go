package previewawsemrcontainersmixins


// The information about the Amazon EKS cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   eksInfoProperty := &EksInfoProperty{
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-eksinfo.html
//
type CfnVirtualClusterPropsMixin_EksInfoProperty struct {
	// The namespaces of the EKS cluster.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 63
	//
	// *Pattern* : `[a-z0-9]([-a-z0-9]*[a-z0-9])?`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-virtualcluster-eksinfo.html#cfn-emrcontainers-virtualcluster-eksinfo-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

