package awssagemaker


// Properties for defining a `CfnContext`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContextProps := &CfnContextProps{
//   	ContextName: jsii.String("contextName"),
//   	ContextType: jsii.String("contextType"),
//   	Source: &SourceProperty{
//   		SourceUri: jsii.String("sourceUri"),
//
//   		// the properties below are optional
//   		SourceId: jsii.String("sourceId"),
//   		SourceType: jsii.String("sourceType"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
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
type CfnContextProps struct {
	// The name of the context.
	//
	// Must be unique to your account in an AWS Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-contextname
	//
	ContextName *string `field:"required" json:"contextName" yaml:"contextName"`
	// The context type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-contexttype
	//
	ContextType *string `field:"required" json:"contextType" yaml:"contextType"`
	// The source type, ID, and URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
	// The description of the context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of properties to add to the context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// A list of tags to apply to the context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-context.html#cfn-sagemaker-context-tags
	//
	Tags *[]*CfnContext_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

