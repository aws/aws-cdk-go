package awseventschemas


// Properties for defining a `CfnDiscoverer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDiscovererProps := &cfnDiscovererProps{
//   	sourceArn: jsii.String("sourceArn"),
//
//   	// the properties below are optional
//   	crossAccount: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDiscovererProps struct {
	// The ARN of the event bus.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// Allows for the discovery of the event schemas that are sent to the event bus from another account.
	CrossAccount interface{} `field:"optional" json:"crossAccount" yaml:"crossAccount"`
	// A description for the discoverer.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags associated with the resource.
	Tags *[]*CfnDiscoverer_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

