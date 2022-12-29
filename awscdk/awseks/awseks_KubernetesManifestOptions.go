package awseks


// Options for `KubernetesManifest`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kubernetesManifestOptions := &kubernetesManifestOptions{
//   	ingressAlb: jsii.Boolean(false),
//   	ingressAlbScheme: awscdk.Aws_eks.albScheme_INTERNAL,
//   	prune: jsii.Boolean(false),
//   	skipValidation: jsii.Boolean(false),
//   }
//
type KubernetesManifestOptions struct {
	// Automatically detect `Ingress` resources in the manifest and annotate them so they are picked up by an ALB Ingress Controller.
	IngressAlb *bool `field:"optional" json:"ingressAlb" yaml:"ingressAlb"`
	// Specify the ALB scheme that should be applied to `Ingress` resources.
	//
	// Only applicable if `ingressAlb` is set to `true`.
	IngressAlbScheme AlbScheme `field:"optional" json:"ingressAlbScheme" yaml:"ingressAlbScheme"`
	// When a resource is removed from a Kubernetes manifest, it no longer appears in the manifest, and there is no way to know that this resource needs to be deleted.
	//
	// To address this, `kubectl apply` has a `--prune` option which will
	// query the cluster for all resources with a specific label and will remove
	// all the labeld resources that are not part of the applied manifest. If this
	// option is disabled and a resource is removed, it will become "orphaned" and
	// will not be deleted from the cluster.
	//
	// When this option is enabled (default), the construct will inject a label to
	// all Kubernetes resources included in this manifest which will be used to
	// prune resources when the manifest changes via `kubectl apply --prune`.
	//
	// The label name will be `aws.cdk.eks/prune-<ADDR>` where `<ADDR>` is the
	// 42-char unique address of this construct in the construct tree. Value is
	// empty.
	// See: https://kubernetes.io/docs/tasks/manage-kubernetes-objects/declarative-config/#alternative-kubectl-apply-f-directory-prune-l-your-label
	//
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// A flag to signify if the manifest validation should be skipped.
	SkipValidation *bool `field:"optional" json:"skipValidation" yaml:"skipValidation"`
}

