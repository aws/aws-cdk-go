package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Kubectl Provider Attributes.
//
// Example:
//   handlerRole := iam.role.fromRoleArn(this, jsii.String("HandlerRole"), jsii.String("arn:aws:iam::123456789012:role/lambda-role"))
//   kubectlProvider := eks.kubectlProvider.fromKubectlProviderAttributes(this, jsii.String("KubectlProvider"), &kubectlProviderAttributes{
//   	functionArn: jsii.String("arn:aws:lambda:us-east-2:123456789012:function:my-function:1"),
//   	kubectlRoleArn: jsii.String("arn:aws:iam::123456789012:role/kubectl-role"),
//   	handlerRole: handlerRole,
//   })
//
//   cluster := eks.cluster.fromClusterAttributes(this, jsii.String("Cluster"), &clusterAttributes{
//   	clusterName: jsii.String("cluster"),
//   	kubectlProvider: kubectlProvider,
//   })
//
// Experimental.
type KubectlProviderAttributes struct {
	// The kubectl provider lambda arn.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The IAM execution role of the handler.
	//
	// This role must be able to assume kubectlRoleArn.
	// Experimental.
	HandlerRole awsiam.IRole `field:"required" json:"handlerRole" yaml:"handlerRole"`
	// The IAM role to assume in order to perform kubectl operations against this cluster.
	// Experimental.
	KubectlRoleArn *string `field:"required" json:"kubectlRoleArn" yaml:"kubectlRoleArn"`
}

