package awsquicksight


// The configuration of the `Select all` options in a list control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listControlSelectAllOptionsProperty := &ListControlSelectAllOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_ListControlSelectAllOptionsProperty struct {
	// The visibility configuration of the `Select all` options in a list control.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

