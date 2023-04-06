package awsquicksight


// The configuration of the search options in a list control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listControlSearchOptionsProperty := &ListControlSearchOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   }
//
type CfnDashboard_ListControlSearchOptionsProperty struct {
	// The visibility configuration of the search options in a list control.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

