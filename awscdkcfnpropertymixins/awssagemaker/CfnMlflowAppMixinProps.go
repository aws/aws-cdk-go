package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMlflowAppPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMlflowAppMixinProps := &CfnMlflowAppMixinProps{
//   	ArtifactStoreUri: jsii.String("artifactStoreUri"),
//   	ModelRegistrationMode: jsii.String("modelRegistrationMode"),
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WeeklyMaintenanceWindowStart: jsii.String("weeklyMaintenanceWindowStart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowapp.html
//
type CfnMlflowAppMixinProps struct {
	// The S3 URI for a general purpose bucket to use as the MLflow App artifact store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowapp.html#cfn-sagemaker-mlflowapp-artifactstoreuri
	//
	ArtifactStoreUri *string `field:"optional" json:"artifactStoreUri" yaml:"artifactStoreUri"`
	// Whether to enable or disable automatic registration of new MLflow models to the SageMaker Model Registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowapp.html#cfn-sagemaker-mlflowapp-modelregistrationmode
	//
	ModelRegistrationMode *string `field:"optional" json:"modelRegistrationMode" yaml:"modelRegistrationMode"`
	// The name of the MLflow App.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowapp.html#cfn-sagemaker-mlflowapp-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) for an IAM role in your account that the MLflow App uses to access the artifact store in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowapp.html#cfn-sagemaker-mlflowapp-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Tags to associate with the MLflow App.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowapp.html#cfn-sagemaker-mlflowapp-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The day and time of the week in Coordinated Universal Time (UTC) 24-hour standard time that weekly maintenance updates are scheduled.
	//
	// For example: Tue:03:30.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowapp.html#cfn-sagemaker-mlflowapp-weeklymaintenancewindowstart
	//
	WeeklyMaintenanceWindowStart *string `field:"optional" json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
}

