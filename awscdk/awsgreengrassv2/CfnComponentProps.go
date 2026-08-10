package awsgreengrassv2


// Properties for defining a `CfnComponent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComponentProps := &CfnComponentProps{
//   	ComponentName: jsii.String("componentName"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-component.html
//
type CfnComponentProps struct {
	// The name of the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-component.html#cfn-greengrassv2-component-componentname
	//
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// Tags associated with the component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-component.html#cfn-greengrassv2-component-tags
	//
	Tags *[]*CfnComponent_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

