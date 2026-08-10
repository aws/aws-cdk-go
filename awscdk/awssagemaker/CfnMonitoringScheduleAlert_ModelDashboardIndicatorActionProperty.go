package awssagemaker


// An alert action taken to light up an icon on the Amazon SageMaker Model Dashboard when an alert goes into InAlert status.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelDashboardIndicatorActionProperty := &ModelDashboardIndicatorActionProperty{
//   	Enabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedulealert-modeldashboardindicatoraction.html
//
type CfnMonitoringScheduleAlert_ModelDashboardIndicatorActionProperty struct {
	// Indicates whether the alert action is turned on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-monitoringschedulealert-modeldashboardindicatoraction.html#cfn-sagemaker-monitoringschedulealert-modeldashboardindicatoraction-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

