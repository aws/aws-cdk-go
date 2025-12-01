package previewawseksmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCapabilityPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityMixinProps := &CfnCapabilityMixinProps{
//   	CapabilityName: jsii.String("capabilityName"),
//   	ClusterName: jsii.String("clusterName"),
//   	Configuration: &CapabilityConfigurationProperty{
//   		ArgoCd: &ArgoCdProperty{
//   			AwsIdc: &AwsIdcProperty{
//   				IdcInstanceArn: jsii.String("idcInstanceArn"),
//   				IdcManagedApplicationArn: jsii.String("idcManagedApplicationArn"),
//   				IdcRegion: jsii.String("idcRegion"),
//   			},
//   			Namespace: jsii.String("namespace"),
//   			NetworkAccess: &NetworkAccessProperty{
//   				VpceIds: []*string{
//   					jsii.String("vpceIds"),
//   				},
//   			},
//   			RbacRoleMappings: []interface{}{
//   				&ArgoCdRoleMappingProperty{
//   					Identities: []interface{}{
//   						&SsoIdentityProperty{
//   							Id: jsii.String("id"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					Role: jsii.String("role"),
//   				},
//   			},
//   			ServerUrl: jsii.String("serverUrl"),
//   		},
//   	},
//   	DeletePropagationPolicy: jsii.String("deletePropagationPolicy"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html
//
type CfnCapabilityMixinProps struct {
	// A unique name for the capability.
	//
	// The name must be unique within your cluster and can contain alphanumeric characters, hyphens, and underscores.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-capabilityname
	//
	CapabilityName *string `field:"optional" json:"capabilityName" yaml:"capabilityName"`
	// The name of the EKS cluster where you want to create the capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-clustername
	//
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// Configuration settings for a capability.
	//
	// The structure of this object varies depending on the capability type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Specifies how Kubernetes resources managed by the capability should be handled when the capability is deleted.
	//
	// Currently, the only supported value is RETAIN which retains all Kubernetes resources managed by the capability when the capability is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-deletepropagationpolicy
	//
	DeletePropagationPolicy *string `field:"optional" json:"deletePropagationPolicy" yaml:"deletePropagationPolicy"`
	// The Amazon Resource Name (ARN) of the IAM role that the capability uses to interact with AWS services.
	//
	// This role must have a trust policy that allows the EKS service principal to assume it, and it must have the necessary permissions for the capability type you're creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of capability to create.
	//
	// Valid values are: ACK (AWS Controllers for Kubernetes, which lets you manage AWS resources directly from Kubernetes), ARGOCD (Argo CD for GitOps-based continuous delivery), or KRO (Kube Resource Orchestrator for composing and managing custom Kubernetes resources).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

