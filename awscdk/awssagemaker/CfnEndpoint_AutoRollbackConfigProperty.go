package awssagemaker


// Automatic rollback configuration for handling endpoint deployment failures and recovery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoRollbackConfigProperty := &AutoRollbackConfigProperty{
//   	Alarms: []interface{}{
//   		&AlarmProperty{
//   			AlarmName: jsii.String("alarmName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-autorollbackconfig.html
//
type CfnEndpoint_AutoRollbackConfigProperty struct {
	// List of CloudWatch alarms in your account that are configured to monitor metrics on an endpoint.
	//
	// If any alarms are tripped during a deployment, SageMaker rolls back the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-autorollbackconfig.html#cfn-sagemaker-endpoint-autorollbackconfig-alarms
	//
	Alarms interface{} `field:"required" json:"alarms" yaml:"alarms"`
}

