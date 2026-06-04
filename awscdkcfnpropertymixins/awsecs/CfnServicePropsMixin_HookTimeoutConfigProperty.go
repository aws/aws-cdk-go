package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   hookTimeoutConfigProperty := &HookTimeoutConfigProperty{
//   	Action: jsii.String("action"),
//   	TimeoutInMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-hooktimeoutconfig.html
//
type CfnServicePropsMixin_HookTimeoutConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-hooktimeoutconfig.html#cfn-ecs-service-hooktimeoutconfig-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-hooktimeoutconfig.html#cfn-ecs-service-hooktimeoutconfig-timeoutinminutes
	//
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
}

