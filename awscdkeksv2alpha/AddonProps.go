package awscdkeksv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for creating an Amazon EKS Add-On.
//
// Example:
//   var cluster Cluster
//
//
//   eks.NewAddon(this, jsii.String("Addon"), &AddonProps{
//   	Cluster: Cluster,
//   	AddonName: jsii.String("coredns"),
//   	AddonVersion: jsii.String("v1.11.4-eksbuild.2"),
//   	// whether to preserve the add-on software on your cluster but Amazon EKS stops managing any settings for the add-on.
//   	PreserveOnDelete: jsii.Boolean(false),
//   	ConfigurationValues: map[string]interface{}{
//   		"replicaCount": jsii.Number(2),
//   	},
//   })
//
// Experimental.
type AddonProps struct {
	// Name of the Add-On.
	// Experimental.
	AddonName *string `field:"required" json:"addonName" yaml:"addonName"`
	// The EKS cluster the Add-On is associated with.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// Version of the Add-On.
	//
	// You can check all available versions with describe-addon-versions.
	// For example, this lists all available versions for the `eks-pod-identity-agent` addon:
	// $ aws eks describe-addon-versions --addon-name eks-pod-identity-agent \
	// --query 'addons[*].addonVersions[*].addonVersion'
	// Default: the latest version.
	//
	// Experimental.
	AddonVersion *string `field:"optional" json:"addonVersion" yaml:"addonVersion"`
	// The configuration values for the Add-on.
	// Default: - Use default configuration.
	//
	// Experimental.
	ConfigurationValues *map[string]interface{} `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	// Specifying this option preserves the add-on software on your cluster but Amazon EKS stops managing any settings for the add-on.
	//
	// If an IAM account is associated with the add-on, it isn't removed.
	// Default: true.
	//
	// Experimental.
	PreserveOnDelete *bool `field:"optional" json:"preserveOnDelete" yaml:"preserveOnDelete"`
	// The removal policy applied to the EKS add-on.
	//
	// The removal policy controls what happens to the resource if it stops being managed by CloudFormation.
	// This can happen in one of three situations:
	//
	// - The resource is removed from the template, so CloudFormation stops managing it
	// - A change to the resource is made that requires it to be replaced, so CloudFormation stops managing it
	// - The stack is deleted, so CloudFormation stops managing all resources in it.
	// Default: RemovalPolicy.DESTROY
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

