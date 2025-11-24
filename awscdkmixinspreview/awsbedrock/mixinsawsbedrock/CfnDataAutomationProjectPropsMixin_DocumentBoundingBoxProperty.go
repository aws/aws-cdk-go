package mixinsawsbedrock


// Bounding box settings for documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   documentBoundingBoxProperty := &DocumentBoundingBoxProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentboundingbox.html
//
type CfnDataAutomationProjectPropsMixin_DocumentBoundingBoxProperty struct {
	// Whether bounding boxes are enabled for documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentboundingbox.html#cfn-bedrock-dataautomationproject-documentboundingbox-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

