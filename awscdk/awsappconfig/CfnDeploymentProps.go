package awsappconfig

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDeployment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentProps := &CfnDeploymentProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	ConfigurationProfileId: jsii.String("configurationProfileId"),
//   	ConfigurationVersion: jsii.String("configurationVersion"),
//   	DeploymentStrategyId: jsii.String("deploymentStrategyId"),
//   	EnvironmentId: jsii.String("environmentId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	DynamicExtensionParameters: []interface{}{
//   		&DynamicExtensionParametersProperty{
//   			ExtensionReference: jsii.String("extensionReference"),
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	KmsKeyIdentifier: jsii.String("kmsKeyIdentifier"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html
//
type CfnDeploymentProps struct {
	// The application ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html#cfn-appconfig-deployment-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The configuration profile ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html#cfn-appconfig-deployment-configurationprofileid
	//
	ConfigurationProfileId *string `field:"required" json:"configurationProfileId" yaml:"configurationProfileId"`
	// The configuration version to deploy.
	//
	// If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html#cfn-appconfig-deployment-configurationversion
	//
	ConfigurationVersion *string `field:"required" json:"configurationVersion" yaml:"configurationVersion"`
	// The deployment strategy ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html#cfn-appconfig-deployment-deploymentstrategyid
	//
	DeploymentStrategyId *string `field:"required" json:"deploymentStrategyId" yaml:"deploymentStrategyId"`
	// The environment ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html#cfn-appconfig-deployment-environmentid
	//
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// A description of the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html#cfn-appconfig-deployment-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of dynamic extension parameter names to values to pass to associated extensions with `PRE_START_DEPLOYMENT` actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html#cfn-appconfig-deployment-dynamicextensionparameters
	//
	DynamicExtensionParameters interface{} `field:"optional" json:"dynamicExtensionParameters" yaml:"dynamicExtensionParameters"`
	// The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html#cfn-appconfig-deployment-kmskeyidentifier
	//
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Metadata to assign to the deployment.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-deployment.html#cfn-appconfig-deployment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

