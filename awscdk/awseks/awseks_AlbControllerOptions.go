package awseks


// Options for `AlbController`.
//
// Example:
//   eks.NewCluster(this, jsii.String("HelloEKS"), &clusterProps{
//   	version: eks.kubernetesVersion_V1_21(),
//   	albController: &albControllerOptions{
//   		version: eks.albControllerVersion_V2_4_1(),
//   	},
//   })
//
type AlbControllerOptions struct {
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
}

