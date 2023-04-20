package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetTextBoxProperty := &SheetTextBoxProperty{
//   	SheetTextBoxId: jsii.String("sheetTextBoxId"),
//
//   	// the properties below are optional
//   	Content: jsii.String("content"),
//   }
//
type CfnTemplate_SheetTextBoxProperty struct {
	// `CfnTemplate.SheetTextBoxProperty.SheetTextBoxId`.
	SheetTextBoxId *string `field:"required" json:"sheetTextBoxId" yaml:"sheetTextBoxId"`
	// `CfnTemplate.SheetTextBoxProperty.Content`.
	Content *string `field:"optional" json:"content" yaml:"content"`
}

