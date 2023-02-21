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
//   serviceAccountProps := &ServiceAccountProps{
//   	Cluster: cluster,
//
//   	// the properties below are optional
//   	Annotations: map[string]*string{
//   		"annotationsKey": jsii.String("annotations"),
//   	},
//   	Labels: map[string]*string{
//   		"labelsKey": jsii.String("labels"),
//   	},
//   	Name: jsii.String("name"),
//   	Namespace: jsii.String("namespace"),
//   }
//
type ServiceAccountProps struct {
	// Additional annotations of the service account.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Additional labels of the service account.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the service account.
	//
	// The name of a ServiceAccount object must be a valid DNS subdomain name.
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the service account.
	//
	// All namespace names must be valid RFC 1123 DNS labels.
	// https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#namespaces-and-dns
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The cluster to apply the patch to.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}

