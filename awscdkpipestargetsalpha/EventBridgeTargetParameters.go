package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// EventBridge target properties.
//
// Example:
//   var sourceQueue queue
//   var targetEventBus eventBus
//
//
//   eventBusTarget := targets.NewEventBridgeTarget(targetEventBus, &EventBridgeTargetParameters{
//   	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
//   		"body": jsii.String("👀"),
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: eventBusTarget,
//   })
//
// Experimental.
type EventBridgeTargetParameters struct {
	// A free-form string used to decide what fields to expect in the event detail.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargeteventbridgeeventbusparameters.html#cfn-pipes-pipe-pipetargeteventbridgeeventbusparameters-detailtype
	//
	// Default: - none.
	//
	// Experimental.
	DetailType *string `field:"optional" json:"detailType" yaml:"detailType"`
	// The URL subdomain of the endpoint.
	//
	// Example:
	//   // if the URL for the endpoint is https://abcde.veo.endpoints.event.amazonaws.com
	//   "abcde.veo"
	//
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargeteventbridgeeventbusparameters.html#cfn-pipes-pipe-pipetargeteventbridgeeventbusparameters-endpointid
	//
	// Default: - none.
	//
	// Experimental.
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: - none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
	// AWS resources, identified by Amazon Resource Name (ARN), which the event primarily concerns.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargeteventbridgeeventbusparameters.html#cfn-pipes-pipe-pipetargeteventbridgeeventbusparameters-resources
	//
	// Default: - none.
	//
	// Experimental.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// The source of the event.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargeteventbridgeeventbusparameters.html#cfn-pipes-pipe-pipetargeteventbridgeeventbusparameters-source
	//
	// Default: - none.
	//
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The time stamp of the event, per RFC3339.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargeteventbridgeeventbusparameters.html#cfn-pipes-pipe-pipetargeteventbridgeeventbusparameters-time
	//
	// Default: - the time stamp of the PutEvents call.
	//
	// Experimental.
	Time *string `field:"optional" json:"time" yaml:"time"`
}

