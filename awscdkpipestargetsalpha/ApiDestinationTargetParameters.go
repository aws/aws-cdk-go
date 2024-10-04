package awscdkpipestargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdkpipesalpha/v2"
)

// EventBridge API destination target properties.
//
// Example:
//   var sourceQueue queue
//   var dest apiDestination
//
//
//   apiTarget := targets.NewApiDestinationTarget(dest, &ApiDestinationTargetParameters{
//   	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
//   		"body": jsii.String("👀"),
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: apiTarget,
//   })
//
// Experimental.
type ApiDestinationTargetParameters struct {
	// The headers to send as part of the request invoking the EventBridge API destination.
	//
	// The headers are merged with the headers from the API destination.
	// If there are conflicts, the headers from the API destination take precedence.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html#cfn-pipes-pipe-pipetargethttpparameters-headerparameters
	//
	// Default: - none.
	//
	// Experimental.
	HeaderParameters *map[string]*string `field:"optional" json:"headerParameters" yaml:"headerParameters"`
	// The input transformation to apply to the message before sending it to the target.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetparameters.html#cfn-pipes-pipe-pipetargetparameters-inputtemplate
	//
	// Default: - none.
	//
	// Experimental.
	InputTransformation awscdkpipesalpha.IInputTransformation `field:"optional" json:"inputTransformation" yaml:"inputTransformation"`
	// The path parameter values used to populate the EventBridge API destination path wildcards ("*").
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html#cfn-pipes-pipe-pipetargethttpparameters-pathparametervalues
	//
	// Default: - none.
	//
	// Experimental.
	PathParameterValues *[]*string `field:"optional" json:"pathParameterValues" yaml:"pathParameterValues"`
	// The query string keys/values that need to be sent as part of request invoking the EventBridge API destination.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargethttpparameters.html#cfn-pipes-pipe-pipetargethttpparameters-querystringparameters
	//
	// Default: - none.
	//
	// Experimental.
	QueryStringParameters *map[string]*string `field:"optional" json:"queryStringParameters" yaml:"queryStringParameters"`
}

