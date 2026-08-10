package awssagemaker


// Properties for CfnContextPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnContextMixinProps := &CfnContextMixinProps{
//   	ContextName: jsii.String("contextName"),
//   	ContextType: jsii.String("contextType"),
//   	Description: jsii.String("description"),
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   	Source: &SourceProperty{
//   		SourceId: jsii.String("sourceId"),
//   		SourceType: jsii.String("sourceType"),
//   		SourceUri: jsii.String("sourceUri"),
//   	},
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html
//
type CfnContextMixinProps struct {
	// The name of the context.
	//
	// Must be unique to your account in an AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-contextname
	//
	ContextName *string `field:"optional" json:"contextName" yaml:"contextName"`
	// The context type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-contexttype
	//
	ContextType *string `field:"optional" json:"contextType" yaml:"contextType"`
	// The description of the context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of properties to add to the context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The source type, ID, and URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// A list of tags to apply to the context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-tags
	//
	Tags *[]*CfnContextPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

