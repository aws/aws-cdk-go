package awscdkeksv2alpha


// Properties for defining service accounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeksv2alpha"
//
//   var cluster Cluster
//
//   serviceAccountProps := &ServiceAccountProps{
//   	Cluster: cluster,
//
//   	// the properties below are optional
//   	Annotations: map[string]*string{
//   		"annotationsKey": jsii.String("annotations"),
//   	},
//   	IdentityType: eks_v2_alpha.IdentityType_IRSA,
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	Name: jsii.String("name"),
//   	Namespace: jsii.String("namespace"),
//   	OverwriteServiceAccount: jsii.Boolean(false),
//   }
//
// Experimental.
type ServiceAccountProps struct {
	// Additional annotations of the service account.
	// Default: - no additional annotations.
	//
	// Experimental.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// The identity type to use for the service account.
	// Default: IdentityType.IRSA
	//
	// Experimental.
	IdentityType IdentityType `field:"optional" json:"identityType" yaml:"identityType"`
	// Additional labels of the service account.
	// Default: - no additional labels.
	//
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the service account.
	//
	// The name of a ServiceAccount object must be a valid DNS subdomain name.
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// Default: - If no name is given, it will use the id of the resource.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the service account.
	//
	// All namespace names must be valid RFC 1123 DNS labels.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#namespaces-and-dns
	// Default: "default".
	//
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Overwrite existing service account.
	//
	// If this is set, we will use `kubectl apply` instead of `kubectl create`
	// when the service account is created. Otherwise, if there is already a service account
	// in the cluster with the same name, the operation will fail.
	// Default: false.
	//
	// Experimental.
	OverwriteServiceAccount *bool `field:"optional" json:"overwriteServiceAccount" yaml:"overwriteServiceAccount"`
	// The cluster to apply the patch to.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

