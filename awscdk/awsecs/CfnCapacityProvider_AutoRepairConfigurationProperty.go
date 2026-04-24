package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoRepairConfigurationProperty := &AutoRepairConfigurationProperty{
//   	ActionsStatus: jsii.String("actionsStatus"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autorepairconfiguration.html
//
type CfnCapacityProvider_AutoRepairConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-autorepairconfiguration.html#cfn-ecs-capacityprovider-autorepairconfiguration-actionsstatus
	//
	ActionsStatus *string `field:"optional" json:"actionsStatus" yaml:"actionsStatus"`
}

