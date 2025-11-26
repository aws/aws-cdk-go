package previewawslaunchwizardmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDeploymentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeploymentMixinProps := &CfnDeploymentMixinProps{
//   	DeploymentPatternName: jsii.String("deploymentPatternName"),
//   	Name: jsii.String("name"),
//   	Specifications: map[string]*string{
//   		"specificationsKey": jsii.String("specifications"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkloadName: jsii.String("workloadName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-launchwizard-deployment.html
//
type CfnDeploymentMixinProps struct {
	// The name of the deployment pattern.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-launchwizard-deployment.html#cfn-launchwizard-deployment-deploymentpatternname
	//
	DeploymentPatternName *string `field:"optional" json:"deploymentPatternName" yaml:"deploymentPatternName"`
	// The name of the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-launchwizard-deployment.html#cfn-launchwizard-deployment-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The settings specified for the deployment.
	//
	// These settings define how to deploy and configure your resources created by the deployment. For more information about the specifications required for creating a deployment for a SAP workload, see [SAP deployment specifications](https://docs.aws.amazon.com/launchwizard/latest/APIReference/launch-wizard-specifications-sap.html) . To retrieve the specifications required to create a deployment for other workloads, use the [`GetWorkloadDeploymentPattern`](https://docs.aws.amazon.com/launchwizard/latest/APIReference/API_GetWorkloadDeploymentPattern.html) operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-launchwizard-deployment.html#cfn-launchwizard-deployment-specifications
	//
	Specifications interface{} `field:"optional" json:"specifications" yaml:"specifications"`
	// Information about the tags attached to a deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-launchwizard-deployment.html#cfn-launchwizard-deployment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-launchwizard-deployment.html#cfn-launchwizard-deployment-workloadname
	//
	WorkloadName *string `field:"optional" json:"workloadName" yaml:"workloadName"`
}

