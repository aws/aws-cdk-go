package awsbedrock


// Bounding box settings for documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentBoundingBoxProperty := &DocumentBoundingBoxProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentboundingbox.html
//
type CfnDataAutomationProject_DocumentBoundingBoxProperty struct {
	// Whether bounding boxes are enabled for documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentboundingbox.html#cfn-bedrock-dataautomationproject-documentboundingbox-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

