package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a KubectlProvider.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster Cluster
//
//   kubectlProviderProps := &KubectlProviderProps{
//   	Cluster: cluster,
//
//   	// the properties below are optional
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   }
//
type KubectlProviderProps struct {
	// The cluster to control.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The removal policy applied to the custom resource that provides kubectl.
	//
	// The removal policy controls what happens to the resource if it stops being managed by CloudFormation.
	// This can happen in one of three situations:
	//
	// - The resource is removed from the template, so CloudFormation stops managing it
	// - A change to the resource is made that requires it to be replaced, so CloudFormation stops managing it
	// - The stack is deleted, so CloudFormation stops managing all resources in it.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

