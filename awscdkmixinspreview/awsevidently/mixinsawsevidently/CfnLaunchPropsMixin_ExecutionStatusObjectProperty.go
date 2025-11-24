package mixinsawsevidently


// Use this structure to start and stop the launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   executionStatusObjectProperty := &ExecutionStatusObjectProperty{
//   	DesiredState: jsii.String("desiredState"),
//   	Reason: jsii.String("reason"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-executionstatusobject.html
//
type CfnLaunchPropsMixin_ExecutionStatusObjectProperty struct {
	// If you are using CloudFormation to stop this launch, specify either `COMPLETED` or `CANCELLED` here to indicate how to classify this experiment.
	//
	// If you omit this parameter, the default of `COMPLETED` is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-executionstatusobject.html#cfn-evidently-launch-executionstatusobject-desiredstate
	//
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// If you are using CloudFormation to stop this launch, this is an optional field that you can use to record why the launch is being stopped or cancelled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-executionstatusobject.html#cfn-evidently-launch-executionstatusobject-reason
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// To start the launch now, specify `START` for this parameter.
	//
	// If this launch is currently running and you want to stop it now, specify `STOP` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-evidently-launch-executionstatusobject.html#cfn-evidently-launch-executionstatusobject-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

