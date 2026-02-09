package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for `AlbController`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var albControllerVersion AlbControllerVersion
//   var cluster Cluster
//   var policy interface{}
//
//   albControllerProps := &AlbControllerProps{
//   	Cluster: cluster,
//   	Version: albControllerVersion,
//
//   	// the properties below are optional
//   	AdditionalHelmChartValues: &AlbControllerHelmChartOptions{
//   		EnableWaf: jsii.Boolean(false),
//   		EnableWafv2: jsii.Boolean(false),
//   	},
//   	OverwriteServiceAccount: jsii.Boolean(false),
//   	Policy: policy,
//   	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
//   	Repository: jsii.String("repository"),
//   }
//
type AlbControllerProps struct {
	// Version of the controller.
	Version AlbControllerVersion `field:"required" json:"version" yaml:"version"`
	// Additional helm chart values for ALB controller.
	// Default: - no additional helm chart values.
	//
	AdditionalHelmChartValues *AlbControllerHelmChartOptions `field:"optional" json:"additionalHelmChartValues" yaml:"additionalHelmChartValues"`
	// Overwrite any existing ALB controller service account.
	//
	// If this is set, we will use `kubectl apply` instead of `kubectl create`
	// when the ALB controller service account is created. Otherwise, if there is already a service account
	// named 'aws-load-balancer-controller' in the kube-system namespace, the operation will fail.
	// Default: false.
	//
	OverwriteServiceAccount *bool `field:"optional" json:"overwriteServiceAccount" yaml:"overwriteServiceAccount"`
	// The IAM policy to apply to the service account.
	//
	// If you're using one of the built-in versions, this is not required since
	// CDK ships with the appropriate policies for those versions.
	//
	// However, if you are using a custom version, this is required (and validated).
	// Default: - Corresponds to the predefined version.
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The removal policy applied to the ALB controller resources.
	//
	// The removal policy controls what happens to the resources if they stop being managed by CloudFormation.
	// This can happen in one of three situations:
	//
	// - The resource is removed from the template, so CloudFormation stops managing it
	// - A change to the resource is made that requires it to be replaced, so CloudFormation stops managing it
	// - The stack is deleted, so CloudFormation stops managing all resources in it.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The repository to pull the controller image from.
	//
	// Note that the default repository works for most regions, but not all.
	// If the repository is not applicable to your region, use a custom repository
	// according to the information here: https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases.
	// Default: '602401143452.dkr.ecr.us-west-2.amazonaws.com/amazon/aws-load-balancer-controller'
	//
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// [disable-awslint:ref-via-interface] Cluster to install the controller onto.
	Cluster Cluster `field:"required" json:"cluster" yaml:"cluster"`
}

