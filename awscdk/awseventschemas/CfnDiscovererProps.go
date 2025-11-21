package awseventschemas


// Properties for defining a `CfnDiscoverer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDiscovererProps := &CfnDiscovererProps{
//   	SourceArn: jsii.String("sourceArn"),
//
//   	// the properties below are optional
//   	CrossAccount: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	Tags: []TagsEntryProperty{
//   		&TagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html
//
type CfnDiscovererProps struct {
	// The ARN of the event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-sourcearn
	//
	SourceArn interface{} `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// Allows for the discovery of the event schemas that are sent to the event bus from another account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-crossaccount
	//
	// Default: - true.
	//
	CrossAccount interface{} `field:"optional" json:"crossAccount" yaml:"crossAccount"`
	// A description for the discoverer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags associated with the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-tags
	//
	Tags *[]*CfnDiscoverer_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

