package awscdkeksv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents the properties required to create an Amazon EKS access entry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPolicy AccessPolicy
//   var cluster Cluster
//
//   accessEntryProps := &AccessEntryProps{
//   	AccessPolicies: []IAccessPolicy{
//   		accessPolicy,
//   	},
//   	Cluster: cluster,
//   	Principal: jsii.String("principal"),
//
//   	// the properties below are optional
//   	AccessEntryName: jsii.String("accessEntryName"),
//   	AccessEntryType: eks_v2_alpha.AccessEntryType_STANDARD,
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   }
//
// Experimental.
type AccessEntryProps struct {
	// The access policies that define the permissions and scope for the access entry.
	// Experimental.
	AccessPolicies *[]IAccessPolicy `field:"required" json:"accessPolicies" yaml:"accessPolicies"`
	// The Amazon EKS cluster to which the access entry applies.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The Amazon Resource Name (ARN) of the principal (user or role) to associate the access entry with.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
	// The name of the AccessEntry.
	// Default: - No access entry name is provided.
	//
	// Experimental.
	AccessEntryName *string `field:"optional" json:"accessEntryName" yaml:"accessEntryName"`
	// The type of the AccessEntry.
	// Default: STANDARD.
	//
	// Experimental.
	AccessEntryType AccessEntryType `field:"optional" json:"accessEntryType" yaml:"accessEntryType"`
	// The removal policy applied to the access entry.
	//
	// The removal policy controls what happens to the resources if they stop being managed by CloudFormation.
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

