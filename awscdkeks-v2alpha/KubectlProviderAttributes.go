package awscdkeks-v2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Kubectl Provider Attributes.
//
// Example:
//   import "github.com/cdklabs/awscdk-kubectl-go/kubectlv32"
//
//
//   handlerRole := iam.Role_FromRoleArn(this, jsii.String("HandlerRole"), jsii.String("arn:aws:iam::123456789012:role/lambda-role"))
//   // get the serivceToken from the custom resource provider
//   functionArn := lambda.Function_FromFunctionName(this, jsii.String("ProviderOnEventFunc"), jsii.String("ProviderframeworkonEvent-XXX")).FunctionArn
//   kubectlProvider := eks.KubectlProvider_FromKubectlProviderAttributes(this, jsii.String("KubectlProvider"), &KubectlProviderAttributes{
//   	ServiceToken: functionArn,
//   	Role: handlerRole,
//   })
//
//   cluster := eks.Cluster_FromClusterAttributes(this, jsii.String("Cluster"), &ClusterAttributes{
//   	ClusterName: jsii.String("cluster"),
//   	KubectlProvider: KubectlProvider,
//   })
//
// Experimental.
type KubectlProviderAttributes struct {
	// The kubectl provider lambda arn.
	// Experimental.
	ServiceToken *string `field:"required" json:"serviceToken" yaml:"serviceToken"`
	// The role of the provider lambda function.
	//
	// Only required if you deploy helm charts using this imported provider.
	// Default: - no role.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

