package previewawseventschemasmixins


// Properties for CfnDiscovererPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDiscovererMixinProps := &CfnDiscovererMixinProps{
//   	CrossAccount: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	SourceArn: jsii.String("sourceArn"),
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
type CfnDiscovererMixinProps struct {
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
	// The ARN of the event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-sourcearn
	//
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// Tags associated with the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventschemas-discoverer.html#cfn-eventschemas-discoverer-tags
	//
	Tags *[]*CfnDiscovererPropsMixin_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

