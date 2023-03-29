package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scrollBarOptionsProperty := &ScrollBarOptionsProperty{
//   	Visibility: jsii.String("visibility"),
//   	VisibleRange: &VisibleRangeOptionsProperty{
//   		PercentRange: &PercentVisibleRangeProperty{
//   			From: jsii.Number(123),
//   			To: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnDashboard_ScrollBarOptionsProperty struct {
	// `CfnDashboard.ScrollBarOptionsProperty.Visibility`.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// `CfnDashboard.ScrollBarOptionsProperty.VisibleRange`.
	VisibleRange interface{} `field:"optional" json:"visibleRange" yaml:"visibleRange"`
}

