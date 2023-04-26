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
//   cfnPipelineProps := &CfnPipelineProps{
//   	PipelineActivities: []interface{}{
//   		&ActivityProperty{
//   			AddAttributes: &AddAttributesProperty{
//   				Attributes: map[string]*string{
//   					"attributesKey": jsii.String("attributes"),
//   				},
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Next: jsii.String("next"),
//   			},
//   			Channel: &ChannelProperty{
//   				ChannelName: jsii.String("channelName"),
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Next: jsii.String("next"),
//   			},
//   			Datastore: &DatastoreProperty{
//   				DatastoreName: jsii.String("datastoreName"),
//   				Name: jsii.String("name"),
//   			},
//   			DeviceRegistryEnrich: &DeviceRegistryEnrichProperty{
//   				Attribute: jsii.String("attribute"),
//   				Name: jsii.String("name"),
//   				RoleArn: jsii.String("roleArn"),
//   				ThingName: jsii.String("thingName"),
//
//   				// the properties below are optional
//   				Next: jsii.String("next"),
//   			},
//   			DeviceShadowEnrich: &DeviceShadowEnrichProperty{
//   				Attribute: jsii.String("attribute"),
//   				Name: jsii.String("name"),
//   				RoleArn: jsii.String("roleArn"),
//   				ThingName: jsii.String("thingName"),
//
//   				// the properties below are optional
//   				Next: jsii.String("next"),
//   			},
//   			Filter: &FilterProperty{
//   				Filter: jsii.String("filter"),
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Next: jsii.String("next"),
//   			},
//   			Lambda: &LambdaProperty{
//   				BatchSize: jsii.Number(123),
//   				LambdaName: jsii.String("lambdaName"),
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Next: jsii.String("next"),
//   			},
//   			Math: &MathProperty{
//   				Attribute: jsii.String("attribute"),
//   				Math: jsii.String("math"),
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Next: jsii.String("next"),
//   			},
//   			RemoveAttributes: &RemoveAttributesProperty{
//   				Attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Next: jsii.String("next"),
//   			},
//   			SelectAttributes: &SelectAttributesProperty{
//   				Attributes: []*string{
//   					jsii.String("attributes"),
//   				},
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Next: jsii.String("next"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	PipelineName: jsii.String("pipelineName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

