package awsevidently


// Use this structure to start and stop the launch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   executionStatusObjectProperty := &executionStatusObjectProperty{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	desiredState: jsii.String("desiredState"),
//   	reason: jsii.String("reason"),
//   }
//
type CfnLaunch_ExecutionStatusObjectProperty struct {
	// To start the launch now, specify `START` for this parameter.
	//
	// If this launch is currently running and you want to stop it now, specify `STOP` .
	Status *string `field:"required" json:"status" yaml:"status"`
	// If you are using AWS CloudFormation to stop this launch, specify either `COMPLETED` or `CANCELLED` here to indicate how to classify this experiment.
	//
	// If you omit this parameter, the default of `COMPLETED` is used.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// If you are using AWS CloudFormation to stop this launch, this is an optional field that you can use to record why the launch is being stopped or cancelled.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

