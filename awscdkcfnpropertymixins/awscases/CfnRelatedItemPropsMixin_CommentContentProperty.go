package awscases


// Represents a comment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   commentContentProperty := &CommentContentProperty{
//   	Body: jsii.String("body"),
//   	ContentType: jsii.String("contentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-relateditem-commentcontent.html
//
type CfnRelatedItemPropsMixin_CommentContentProperty struct {
	// Text in the body of a comment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-relateditem-commentcontent.html#cfn-cases-relateditem-commentcontent-body
	//
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Type of the text in the comment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cases-relateditem-commentcontent.html#cfn-cases-relateditem-commentcontent-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

