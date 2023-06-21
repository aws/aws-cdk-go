package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Kubectl Provider Attributes.
//
// Example:
//   handlerRole := iam.Role_FromRoleArn(this, jsii.String("HandlerRole"), jsii.String("arn:aws:iam::123456789012:role/lambda-role"))
//   // get the serivceToken from the custom resource provider
//   functionArn := lambda.Function_FromFunctionName(this, jsii.String("ProviderOnEventFunc"), jsii.String("ProviderframeworkonEvent-XXX")).FunctionArn
//   kubectlProvider := eks.KubectlProvider_FromKubectlProviderAttributes(this, jsii.String("KubectlProvider"), &KubectlProviderAttributes{
//   	FunctionArn: jsii.String(FunctionArn),
//   	KubectlRoleArn: jsii.String("arn:aws:iam::123456789012:role/kubectl-role"),
//   	HandlerRole: HandlerRole,
//   })
//
//   cluster := eks.Cluster_FromClusterAttributes(this, jsii.String("Cluster"), &ClusterAttributes{
//   	ClusterName: jsii.String("cluster"),
//   	KubectlProvider: KubectlProvider,
//   })
//
type KubectlProviderAttributes struct {
	// The custom resource provider's service token.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// The IAM execution role of the handler.
	//
	// This role must be able to assume kubectlRoleArn.
	HandlerRole awsiam.IRole `field:"required" json:"handlerRole" yaml:"handlerRole"`
	// The IAM role to assume in order to perform kubectl operations against this cluster.
	KubectlRoleArn *string `field:"required" json:"kubectlRoleArn" yaml:"kubectlRoleArn"`
}

