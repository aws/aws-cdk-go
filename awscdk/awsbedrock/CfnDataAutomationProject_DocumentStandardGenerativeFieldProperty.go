package awsbedrock


// Settings for generating descriptions of documents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentStandardGenerativeFieldProperty := &DocumentStandardGenerativeFieldProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentstandardgenerativefield.html
//
type CfnDataAutomationProject_DocumentStandardGenerativeFieldProperty struct {
	// Whether generating descriptions is enabled for documents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentstandardgenerativefield.html#cfn-bedrock-dataautomationproject-documentstandardgenerativefield-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

