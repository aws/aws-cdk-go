package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// Amazon Data Firehose target properties.
//
// Example:
//   var sourceQueue queue
//   var targetDeliveryStream deliveryStream
//
//
//   deliveryStreamTarget := targets.NewFirehoseTarget(targetDeliveryStream, &FirehoseTargetParameters{
//   	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
//   		"body": jsii.String("👀"),
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: deliveryStreamTarget,
//   })
//
// Experimental.
type FirehoseTargetParameters struct {
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: - none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
}

