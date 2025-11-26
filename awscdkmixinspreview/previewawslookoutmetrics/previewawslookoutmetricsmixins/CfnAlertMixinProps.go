package previewawslookoutmetricsmixins


// Properties for CfnAlertPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAlertMixinProps := &CfnAlertMixinProps{
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
//   	AlertDescription: jsii.String("alertDescription"),
//   	AlertName: jsii.String("alertName"),
//   	AlertSensitivityThreshold: jsii.Number(123),
//   	AnomalyDetectorArn: jsii.String("anomalyDetectorArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-alert.html
//
type CfnAlertMixinProps struct {
	// Action that will be triggered when there is an alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-alert.html#cfn-lookoutmetrics-alert-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// A description of the alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-alert.html#cfn-lookoutmetrics-alert-alertdescription
	//
	AlertDescription *string `field:"optional" json:"alertDescription" yaml:"alertDescription"`
	// The name of the alert.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-alert.html#cfn-lookoutmetrics-alert-alertname
	//
	AlertName *string `field:"optional" json:"alertName" yaml:"alertName"`
	// An integer from 0 to 100 specifying the alert sensitivity threshold.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-alert.html#cfn-lookoutmetrics-alert-alertsensitivitythreshold
	//
	AlertSensitivityThreshold *float64 `field:"optional" json:"alertSensitivityThreshold" yaml:"alertSensitivityThreshold"`
	// The ARN of the detector to which the alert is attached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lookoutmetrics-alert.html#cfn-lookoutmetrics-alert-anomalydetectorarn
	//
	AnomalyDetectorArn *string `field:"optional" json:"anomalyDetectorArn" yaml:"anomalyDetectorArn"`
}

