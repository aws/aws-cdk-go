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
type CfnDashboard_SheetProperty struct {
	// `CfnDashboard.SheetProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// `CfnDashboard.SheetProperty.SheetId`.
	SheetId *string `field:"optional" json:"sheetId" yaml:"sheetId"`
}

