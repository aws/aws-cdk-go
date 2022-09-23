package awseks


// Properties for `AlbController`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var albControllerVersion albControllerVersion
//   var cluster cluster
//   var policy interface{}
//
//   albControllerProps := &albControllerProps{
//   	cluster: cluster,
//   	version: albControllerVersion,
//
//   	// the properties below are optional
//   	policy: policy,
//   	repository: jsii.String("repository"),
//   }
//
type AlbControllerProps struct {
	// Version of the controller.
	Version AlbControllerVersion `field:"required" json:"version" yaml:"version"`
	// The IAM policy to apply to the service account.
	//
	// If you're using one of the built-in versions, this is not required since
	// CDK ships with the appropriate policies for those versions.
	//
	// However, if you are using a custom version, this is required (and validated).
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The repository to pull the controller image from.
	//
	// Note that the default repository works for most regions, but not all.
	// If the repository is not applicable to your region, use a custom repository
	// according to the information here: https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// [disable-awslint:ref-via-interface] Cluster to install the controller onto.
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
}

