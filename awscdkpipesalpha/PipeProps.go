package awscdkpipesalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for a pipe.
//
// Example:
//   var sourceQueue queue
//   var targetQueue queue
//
//
//   pipeSource := sources.NewSqsSource(sourceQueue, &SqsSourceParameters{
//   	BatchSize: jsii.Number(10),
//   	MaximumBatchingWindow: cdk.Duration_Seconds(jsii.Number(10)),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: pipeSource,
//   	Target: NewSomeTarget(targetQueue),
//   })
//
// Experimental.
type PipeProps struct {
	// The source of the pipe.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-source.html
	//
	// Experimental.
	Source ISource `field:"required" json:"source" yaml:"source"`
	// The target of the pipe.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-target.html
	//
	// Experimental.
	Target ITarget `field:"required" json:"target" yaml:"target"`
	// A description of the pipe displayed in the AWS console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-description
	//
	// Default: - no description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The desired state of the pipe.
	//
	// If the state is set to STOPPED, the pipe will not process events.
	// See: https://docs.aws.amazon.com/eventbridge/latest/pipes-reference/API_Pipe.html#eventbridge-Type-Pipe-DesiredState
	//
	// Default: - DesiredState.RUNNING
	//
	// Experimental.
	DesiredState DesiredState `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Enrichment step to enhance the data from the source before sending it to the target.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/pipes-enrichment.html
	//
	// Default: - no enrichment.
	//
	// Experimental.
	Enrichment IEnrichment `field:"optional" json:"enrichment" yaml:"enrichment"`
	// The filter pattern for the pipe source.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-event-filtering.html
	//
	// Default: - no filter.
	//
	// Experimental.
	Filter IFilter `field:"optional" json:"filter" yaml:"filter"`
	// Destinations for the logs.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-logs.html
	//
	// Default: - no logs.
	//
	// Experimental.
	LogDestinations *[]ILogDestination `field:"optional" json:"logDestinations" yaml:"logDestinations"`
	// Whether the execution data (specifically, the `payload` , `awsRequest` , and `awsResponse` fields) is included in the log messages for this pipe.
	//
	// This applies to all log destinations for the pipe.
	//
	// For more information, see [Including execution data in logs](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-logs.html#eb-pipes-logs-execution-data) and the [message schema](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-logs-schema.html) in the *Amazon EventBridge User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipelogconfiguration.html#cfn-pipes-pipe-pipelogconfiguration-includeexecutiondata
	//
	// Default: - none.
	//
	// Experimental.
	LogIncludeExecutionData *[]IncludeExecutionData `field:"optional" json:"logIncludeExecutionData" yaml:"logIncludeExecutionData"`
	// The level of logging detail to include.
	//
	// This applies to all log destinations for the pipe.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-logs.html
	//
	// Default: - LogLevel.ERROR
	//
	// Experimental.
	LogLevel LogLevel `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Name of the pipe in the AWS console.
	// Default: - automatically generated name.
	//
	// Experimental.
	PipeName *string `field:"optional" json:"pipeName" yaml:"pipeName"`
	// The role used by the pipe which has permissions to read from the source and write to the target.
	//
	// If an enriched target is used, the role also have permissions to call the enriched target.
	// If no role is provided, a role will be created.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-permissions.html
	//
	// Default: - a new role will be created.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The list of key-value pairs to associate with the pipe.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html#cfn-pipes-pipe-tags
	//
	// Default: - no tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

