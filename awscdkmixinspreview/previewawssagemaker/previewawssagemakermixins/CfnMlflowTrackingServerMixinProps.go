package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMlflowTrackingServerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMlflowTrackingServerMixinProps := &CfnMlflowTrackingServerMixinProps{
//   	ArtifactStoreUri: jsii.String("artifactStoreUri"),
//   	AutomaticModelRegistration: jsii.Boolean(false),
//   	MlflowVersion: jsii.String("mlflowVersion"),
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackingServerName: jsii.String("trackingServerName"),
//   	TrackingServerSize: jsii.String("trackingServerSize"),
//   	WeeklyMaintenanceWindowStart: jsii.String("weeklyMaintenanceWindowStart"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowtrackingserver.html
//
type CfnMlflowTrackingServerMixinProps struct {
	// The Amazon S3 URI for MLFlow Tracking Server artifacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowtrackingserver.html#cfn-sagemaker-mlflowtrackingserver-artifactstoreuri
	//
	ArtifactStoreUri *string `field:"optional" json:"artifactStoreUri" yaml:"artifactStoreUri"`
	// A flag to enable Automatic SageMaker Model Registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowtrackingserver.html#cfn-sagemaker-mlflowtrackingserver-automaticmodelregistration
	//
	AutomaticModelRegistration interface{} `field:"optional" json:"automaticModelRegistration" yaml:"automaticModelRegistration"`
	// The MLFlow Version used on the MLFlow Tracking Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowtrackingserver.html#cfn-sagemaker-mlflowtrackingserver-mlflowversion
	//
	MlflowVersion *string `field:"optional" json:"mlflowVersion" yaml:"mlflowVersion"`
	// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowtrackingserver.html#cfn-sagemaker-mlflowtrackingserver-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowtrackingserver.html#cfn-sagemaker-mlflowtrackingserver-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the MLFlow Tracking Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowtrackingserver.html#cfn-sagemaker-mlflowtrackingserver-trackingservername
	//
	TrackingServerName *string `field:"optional" json:"trackingServerName" yaml:"trackingServerName"`
	// The size of the MLFlow Tracking Server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowtrackingserver.html#cfn-sagemaker-mlflowtrackingserver-trackingserversize
	//
	TrackingServerSize *string `field:"optional" json:"trackingServerSize" yaml:"trackingServerSize"`
	// The start of the time window for maintenance of the MLFlow Tracking Server in UTC time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-mlflowtrackingserver.html#cfn-sagemaker-mlflowtrackingserver-weeklymaintenancewindowstart
	//
	WeeklyMaintenanceWindowStart *string `field:"optional" json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
}

