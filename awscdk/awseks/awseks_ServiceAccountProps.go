package awseks


// Properties for defining service accounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//
//   serviceAccountProps := &serviceAccountProps{
//   	cluster: cluster,
//
//   	// the properties below are optional
//   	annotations: map[string]*string{
//   		"annotationsKey": jsii.String("annotations"),
//   	},
//   	labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	name: jsii.String("name"),
//   	namespace: jsii.String("namespace"),
//   }
//
// Experimental.
type ServiceAccountProps struct {
	// Additional annotations of the service account.
	// Experimental.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Additional labels of the service account.
	// Experimental.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the service account.
	//
	// The name of a ServiceAccount object must be a valid DNS subdomain name.
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the service account.
	//
	// All namespace names must be valid RFC 1123 DNS labels.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#namespaces-and-dns
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The cluster to apply the patch to.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

