package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCapability`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCapabilityProps := &CfnCapabilityProps{
//   	CapabilityName: jsii.String("capabilityName"),
//   	ClusterName: jsii.String("clusterName"),
//   	DeletePropagationPolicy: jsii.String("deletePropagationPolicy"),
//   	RoleArn: jsii.String("roleArn"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Configuration: &CapabilityConfigurationProperty{
//   		ArgoCd: &ArgoCdProperty{
//   			AwsIdc: &AwsIdcProperty{
//   				IdcInstanceArn: jsii.String("idcInstanceArn"),
//
//   				// the properties below are optional
//   				IdcManagedApplicationArn: jsii.String("idcManagedApplicationArn"),
//   				IdcRegion: jsii.String("idcRegion"),
//   			},
//
//   			// the properties below are optional
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html
//
type CfnCapabilityProps struct {
	// The unique name of the capability within the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-capabilityname
	//
	CapabilityName *string `field:"required" json:"capabilityName" yaml:"capabilityName"`
	// The name of the Amazon EKS cluster that contains this capability.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-clustername
	//
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The delete propagation policy for the capability.
	//
	// Currently, the only supported value is `RETAIN` , which keeps all resources managed by the capability when the capability is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-deletepropagationpolicy
	//
	DeletePropagationPolicy *string `field:"required" json:"deletePropagationPolicy" yaml:"deletePropagationPolicy"`
	// The Amazon Resource Name (ARN) of the IAM role that the capability uses to interact with AWS services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The type of capability.
	//
	// Valid values are `ACK` , `ARGOCD` , or `KRO` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The configuration settings for the capability.
	//
	// The structure varies depending on the capability type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-capability.html#cfn-eks-capability-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

