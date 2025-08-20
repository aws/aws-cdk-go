package awseks


// Enum representing the different identity types that can be used for a Kubernetes service account.
//
// Example:
//   var cluster cluster
//
//
//   eks.NewServiceAccount(this, jsii.String("ServiceAccount"), &ServiceAccountProps{
//   	Cluster: Cluster,
//   	Name: jsii.String("test-sa"),
//   	Namespace: jsii.String("default"),
//   	IdentityType: eks.IdentityType_POD_IDENTITY,
//   })
//
type IdentityType string

const (
	// Use the IAM Roles for Service Accounts (IRSA) identity type.
	//
	// IRSA allows you to associate an IAM role with a Kubernetes service account.
	// This provides a way to grant permissions to Kubernetes pods by associating an IAM role with a Kubernetes service account.
	// The IAM role can then be used to provide AWS credentials to the pods, allowing them to access other AWS resources.
	//
	// When enabled, the openIdConnectProvider of the cluster would be created when you create the ServiceAccount.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html
	//
	IdentityType_IRSA IdentityType = "IRSA"
	// Use the EKS Pod Identities identity type.
	//
	// EKS Pod Identities provide the ability to manage credentials for your applications, similar to the way that Amazon EC2 instance profiles
	// provide credentials to Amazon EC2 instances. Instead of creating and distributing your AWS credentials to the containers or using the
	// Amazon EC2 instance's role, you associate an IAM role with a Kubernetes service account and configure your Pods to use the service account.
	//
	// When enabled, the Pod Identity Agent AddOn of the cluster would be created when you create the ServiceAccount.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/pod-identities.html
	//
	IdentityType_POD_IDENTITY IdentityType = "POD_IDENTITY"
)

