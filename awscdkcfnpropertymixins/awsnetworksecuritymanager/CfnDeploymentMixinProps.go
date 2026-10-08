package awsnetworksecuritymanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDeploymentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnDeploymentMixinProps := &CfnDeploymentMixinProps{
//   	AssociatedPolicyList: []interface{}{
//   		&AssociatedPolicyProperty{
//   			PolicyArn: jsii.String("policyArn"),
//   		},
//   	},
//   	AssociatedScopeList: []interface{}{
//   		&AssociatedScopeProperty{
//   			ScopeArn: jsii.String("scopeArn"),
//   		},
//   	},
//   	DeploymentConfiguration: &DeploymentConfigurationProperty{
//   		EnableCrossAccountVisibility: jsii.Boolean(false),
//   	},
//   	DeploymentDescription: jsii.String("deploymentDescription"),
//   	DeploymentName: jsii.String("deploymentName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-deployment.html
//
type CfnDeploymentMixinProps struct {
	// List of policies associated with this deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-deployment.html#cfn-networksecuritymanager-deployment-associatedpolicylist
	//
	AssociatedPolicyList interface{} `field:"optional" json:"associatedPolicyList" yaml:"associatedPolicyList"`
	// List of scopes associated with this deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-deployment.html#cfn-networksecuritymanager-deployment-associatedscopelist
	//
	AssociatedScopeList interface{} `field:"optional" json:"associatedScopeList" yaml:"associatedScopeList"`
	// Configuration settings for the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-deployment.html#cfn-networksecuritymanager-deployment-deploymentconfiguration
	//
	DeploymentConfiguration interface{} `field:"optional" json:"deploymentConfiguration" yaml:"deploymentConfiguration"`
	// A description of the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-deployment.html#cfn-networksecuritymanager-deployment-deploymentdescription
	//
	DeploymentDescription *string `field:"optional" json:"deploymentDescription" yaml:"deploymentDescription"`
	// The name of the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-deployment.html#cfn-networksecuritymanager-deployment-deploymentname
	//
	DeploymentName *string `field:"optional" json:"deploymentName" yaml:"deploymentName"`
	// The tags associated with the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networksecuritymanager-deployment.html#cfn-networksecuritymanager-deployment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

