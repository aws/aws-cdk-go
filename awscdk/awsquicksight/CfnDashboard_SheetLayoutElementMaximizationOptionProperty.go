package awsquicksight


// The sheet layout maximization options of a dashbaord.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sheetLayoutElementMaximizationOptionProperty := &SheetLayoutElementMaximizationOptionProperty{
//   	AvailabilityStatus: jsii.String("availabilityStatus"),
//   }
//
type CfnDashboard_SheetLayoutElementMaximizationOptionProperty struct {
	// The status of the sheet layout maximization options of a dashbaord.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
}

