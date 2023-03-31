package awsstepfunctions


// Properties for defining a `CfnActivity`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnActivityProps := &cfnActivityProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnActivityProps struct {
	// The name of the activity.
	//
	// A name must *not* contain:
	//
	// - white space
	// - brackets `< > { } [ ]`
	// - wildcard characters `? *`
	// - special characters `" # % \ ^ | ~ ` $ & , ; : /`
	// - control characters ( `U+0000-001F` , `U+007F-009F` )
	//
	// To enable logging with CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and _.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of tags to add to a resource.
	//
	// Tags may only contain Unicode letters, digits, white space, or these symbols: `_ . : / = + - @` .
	Tags *[]*CfnActivity_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

