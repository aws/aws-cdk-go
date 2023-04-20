package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rowAlternateColorOptionsProperty := &RowAlternateColorOptionsProperty{
//   	RowAlternateColors: []*string{
//   		jsii.String("rowAlternateColors"),
//   	},
//   	Status: jsii.String("status"),
//   }
//
type CfnDashboard_RowAlternateColorOptionsProperty struct {
	// `CfnDashboard.RowAlternateColorOptionsProperty.RowAlternateColors`.
	RowAlternateColors *[]*string `field:"optional" json:"rowAlternateColors" yaml:"rowAlternateColors"`
	// `CfnDashboard.RowAlternateColorOptionsProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

