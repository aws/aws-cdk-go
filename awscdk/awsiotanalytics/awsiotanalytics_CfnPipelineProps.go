package awsiotanalytics

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPipeline`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPipelineProps := &cfnPipelineProps{
//   	pipelineActivities: []interface{}{
//   		&activityProperty{
//   			addAttributes: &addAttributesProperty{
//   				attributes: map[string]*string{
//   					"attributesKey": jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			channel: &channelProperty{
//   				channelName: jsii.String("channelName"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			datastore: &datastoreProperty{
//   				datastoreName: jsii.String("datastoreName"),
//   				name: jsii.String("name"),
//   			},
//   			deviceRegistryEnrich: &deviceRegistryEnrichProperty{
//   				attribute: jsii.String("attribute"),
//   				name: jsii.String("name"),
//   				roleArn: jsii.String("roleArn"),
//   				thingName: jsii.String("thingName"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			deviceShadowEnrich: &deviceShadowEnrichProperty{
//   				attribute: jsii.String("attribute"),
//   				name: jsii.String("name"),
//   				roleArn: jsii.String("roleArn"),
//   				thingName: jsii.String("thingName"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			filter: &filterProperty{
//   				filter: jsii.String("filter"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			lambda: &lambdaProperty{
//   				batchSize: jsii.Number(123),
//   				lambdaName: jsii.String("lambdaName"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			math: &mathProperty{
//   				attribute: jsii.String("attribute"),
//   				math: jsii.String("math"),
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			removeAttributes: &removeAttributesProperty{
//   				attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   			selectAttributes: &selectAttributesProperty{
//   				attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//
//   				// the properties below are optional
//   				next: jsii.String("next"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	pipelineName: jsii.String("pipelineName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPipelineProps struct {
	// A list of "PipelineActivity" objects.
	//
	// Activities perform transformations on your messages, such as removing, renaming or adding message attributes; filtering messages based on attribute values; invoking your Lambda functions on messages for advanced processing; or performing mathematical transformations to normalize device data.
	//
	// The list can be 2-25 *PipelineActivity* objects and must contain both a `channel` and a `datastore` activity. Each entry in the list must contain only one activity, for example:
	//
	// `pipelineActivities = [ { "channel": { ... } }, { "lambda": { ... } }, ... ]`
	PipelineActivities interface{} `field:"required" json:"pipelineActivities" yaml:"pipelineActivities"`
	// The name of the pipeline.
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// Metadata which can be used to manage the pipeline.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

