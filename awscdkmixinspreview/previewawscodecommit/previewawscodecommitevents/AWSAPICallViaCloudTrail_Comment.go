package previewawscodecommitevents


// Type definition for Comment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   comment := &Comment{
//   	AuthorArn: []*string{
//   		jsii.String("authorArn"),
//   	},
//   	ClientRequestToken: []*string{
//   		jsii.String("clientRequestToken"),
//   	},
//   	CommentId: []*string{
//   		jsii.String("commentId"),
//   	},
//   	Content: []*string{
//   		jsii.String("content"),
//   	},
//   	CreationDate: []*string{
//   		jsii.String("creationDate"),
//   	},
//   	Deleted: []*string{
//   		jsii.String("deleted"),
//   	},
//   	InReplyTo: []*string{
//   		jsii.String("inReplyTo"),
//   	},
//   	LastModifiedDate: []*string{
//   		jsii.String("lastModifiedDate"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_Comment struct {
	// authorArn property.
	//
	// Specify an array of string values to match this event if the actual value of authorArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AuthorArn *[]*string `field:"optional" json:"authorArn" yaml:"authorArn"`
	// clientRequestToken property.
	//
	// Specify an array of string values to match this event if the actual value of clientRequestToken is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ClientRequestToken *[]*string `field:"optional" json:"clientRequestToken" yaml:"clientRequestToken"`
	// commentId property.
	//
	// Specify an array of string values to match this event if the actual value of commentId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CommentId *[]*string `field:"optional" json:"commentId" yaml:"commentId"`
	// content property.
	//
	// Specify an array of string values to match this event if the actual value of content is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Content *[]*string `field:"optional" json:"content" yaml:"content"`
	// creationDate property.
	//
	// Specify an array of string values to match this event if the actual value of creationDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreationDate *[]*string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// deleted property.
	//
	// Specify an array of string values to match this event if the actual value of deleted is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Deleted *[]*string `field:"optional" json:"deleted" yaml:"deleted"`
	// inReplyTo property.
	//
	// Specify an array of string values to match this event if the actual value of inReplyTo is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InReplyTo *[]*string `field:"optional" json:"inReplyTo" yaml:"inReplyTo"`
	// lastModifiedDate property.
	//
	// Specify an array of string values to match this event if the actual value of lastModifiedDate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LastModifiedDate *[]*string `field:"optional" json:"lastModifiedDate" yaml:"lastModifiedDate"`
}

