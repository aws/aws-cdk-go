package awsmedialive


// Properties for defining a `CfnCloudWatchAlarmTemplateGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudWatchAlarmTemplateGroupProps := &CfnCloudWatchAlarmTemplateGroupProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplategroup.html
//
type CfnCloudWatchAlarmTemplateGroupProps struct {
	// A resource's name.
	//
	// Names must be unique within the scope of a resource type in a specific region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplategroup.html#cfn-medialive-cloudwatchalarmtemplategroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A resource's optional description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplategroup.html#cfn-medialive-cloudwatchalarmtemplategroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Represents the tags associated with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-cloudwatchalarmtemplategroup.html#cfn-medialive-cloudwatchalarmtemplategroup-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

