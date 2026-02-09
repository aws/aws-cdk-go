package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for KubernetesPatch.
//
// Example:
//   var cluster Cluster
//
//   eks.NewKubernetesPatch(this, jsii.String("hello-kub-deployment-label"), &KubernetesPatchProps{
//   	Cluster: Cluster,
//   	ResourceName: jsii.String("deployment/hello-kubernetes"),
//   	ApplyPatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(5),
//   		},
//   	},
//   	RestorePatch: map[string]interface{}{
//   		"spec": map[string]*f64{
//   			"replicas": jsii.Number(3),
//   		},
//   	},
//   })
//
type KubernetesPatchProps struct {
	// The JSON object to pass to `kubectl patch` when the resource is created/updated.
	ApplyPatch *map[string]interface{} `field:"required" json:"applyPatch" yaml:"applyPatch"`
	// The cluster to apply the patch to.
	//
	// [disable-awslint:ref-via-interface].
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The full name of the resource to patch (e.g. `deployment/coredns`).
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// The JSON object to pass to `kubectl patch` when the resource is removed.
	RestorePatch *map[string]interface{} `field:"required" json:"restorePatch" yaml:"restorePatch"`
	// The patch type to pass to `kubectl patch`.
	//
	// The default type used by `kubectl patch` is "strategic".
	// Default: PatchType.STRATEGIC
	//
	PatchType PatchType `field:"optional" json:"patchType" yaml:"patchType"`
	// The removal policy applied to the custom resource that manages the Kubernetes patch.
	//
	// The removal policy controls what happens to the resource if it stops being managed by CloudFormation.
	// This can happen in one of three situations:
	//
	// - The resource is removed from the template, so CloudFormation stops managing it
	// - A change to the resource is made that requires it to be replaced, so CloudFormation stops managing it
	// - The stack is deleted, so CloudFormation stops managing all resources in it.
	// Default: RemovalPolicy.DESTROY
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The kubernetes API namespace.
	// Default: "default".
	//
	ResourceNamespace *string `field:"optional" json:"resourceNamespace" yaml:"resourceNamespace"`
}

