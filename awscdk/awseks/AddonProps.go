package awseks


// Properties for creating an Amazon EKS Add-On.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   addonProps := &AddonProps{
//   	AddonName: jsii.String("addonName"),
//   	Cluster: cluster,
//
//   	// the properties below are optional
//   	AddonVersion: jsii.String("addonVersion"),
//   }
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
}

