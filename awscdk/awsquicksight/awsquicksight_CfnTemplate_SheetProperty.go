package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetProperty := &sheetProperty{
//   	name: jsii.String("name"),
//   	sheetId: jsii.String("sheetId"),
//   }
//
type CfnTemplate_SheetProperty struct {
	// `CfnTemplate.SheetProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnTemplate.SheetProperty.SheetId`.
	SheetId *string `field:"optional" json:"sheetId" yaml:"sheetId"`
}

