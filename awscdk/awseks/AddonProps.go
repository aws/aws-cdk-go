package awseks


// Properties for creating an Amazon EKS Add-On.
//
// Example:
//   var cluster cluster
//
//
//   eks.NewAddon(this, jsii.String("Addon"), &AddonProps{
//   	Cluster: Cluster,
//   	AddonName: jsii.String("aws-guardduty-agent"),
//   	AddonVersion: jsii.String("v1.6.1"),
//   	// whether to preserve the add-on software on your cluster but Amazon EKS stops managing any settings for the add-on.
//   	PreserveOnDelete: jsii.Boolean(false),
//   })
//
type AddonProps struct {
	// Name of the Add-On.
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// The EKS cluster the Add-On is associated with.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// Version of the Add-On.
	//
	// You can check all available versions with describe-addon-versons.
	// For example, this lists all available versions for the `eks-pod-identity-agent` addon:
	// $ aws eks describe-addon-versions --addon-name eks-pod-identity-agent \
	// --query 'addons[*].addonVersions[*].addonVersion'
	// Default: the latest version.
	//
	AddonVersion *string `field:"optional" json:"addonVersion" yaml:"addonVersion"`
	// Specifying this option preserves the add-on software on your cluster but Amazon EKS stops managing any settings for the add-on.
	//
	// If an IAM account is associated with the add-on, it isn't removed.
	// Default: true.
	//
	PreserveOnDelete *bool `field:"optional" json:"preserveOnDelete" yaml:"preserveOnDelete"`
}

