package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoRollbackConfigurationProperty := &AutoRollbackConfigurationProperty{
//   	Alarms: []interface{}{
//   		&AlarmProperty{
//   			AlarmName: jsii.String("alarmName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-autorollbackconfiguration.html
//
type CfnInferenceComponent_AutoRollbackConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-inferencecomponent-autorollbackconfiguration.html#cfn-sagemaker-inferencecomponent-autorollbackconfiguration-alarms
	//
	Alarms interface{} `field:"required" json:"alarms" yaml:"alarms"`
}

