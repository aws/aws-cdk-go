package awscdkeks-v2alpha


// Properties for `AlbController`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeks-v2alpha"
//
//   var albControllerVersion albControllerVersion
//   var cluster cluster
//   var policy interface{}
//
//   albControllerProps := &AlbControllerProps{
//   	Cluster: cluster,
//   	Version: albControllerVersion,
//
//   	// the properties below are optional
//   	Policy: policy,
//   	Repository: jsii.String("repository"),
//   }
//
// Experimental.
type AlbControllerProps struct {
	// Version of the controller.
	// Experimental.
	Version AlbControllerVersion `field:"required" json:"version" yaml:"version"`
	// The IAM policy to apply to the service account.
	//
	// If you're using one of the built-in versions, this is not required since
	// CDK ships with the appropriate policies for those versions.
	//
	// However, if you are using a custom version, this is required (and validated).
	// Default: - Corresponds to the predefined version.
	//
	// Experimental.
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The repository to pull the controller image from.
	//
	// Note that the default repository works for most regions, but not all.
	// If the repository is not applicable to your region, use a custom repository
	// according to the information here: https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases.
	// Default: '602401143452.dkr.ecr.us-west-2.amazonaws.com/amazon/aws-load-balancer-controller'
	//
	// Experimental.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// [disable-awslint:ref-via-interface] Cluster to install the controller onto.
	// Experimental.
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
}

