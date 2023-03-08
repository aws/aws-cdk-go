package awslookoutmetrics


// Properties for defining a `CfnAlert`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAlertProps := &CfnAlertProps{
//   	Action: &ActionProperty{
//   		LambdaConfiguration: &LambdaConfigurationProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		SnsConfiguration: &SNSConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SnsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	AlertSensitivityThreshold: jsii.Number(123),
//   	AnomalyDetectorArn: jsii.String("anomalyDetectorArn"),
//
//   	// the properties below are optional
//   	AlertDescription: jsii.String("alertDescription"),
//   	AlertName: jsii.String("alertName"),
//   }
//
type CfnAlertProps struct {
	// Action that will be triggered when there is an alert.
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// An integer from 0 to 100 specifying the alert sensitivity threshold.
	AlertSensitivityThreshold *float64 `field:"required" json:"alertSensitivityThreshold" yaml:"alertSensitivityThreshold"`
	// The ARN of the detector to which the alert is attached.
	AnomalyDetectorArn *string `field:"required" json:"anomalyDetectorArn" yaml:"anomalyDetectorArn"`
	// A description of the alert.
	AlertDescription *string `field:"optional" json:"alertDescription" yaml:"alertDescription"`
	// The name of the alert.
	AlertName *string `field:"optional" json:"alertName" yaml:"alertName"`
}

