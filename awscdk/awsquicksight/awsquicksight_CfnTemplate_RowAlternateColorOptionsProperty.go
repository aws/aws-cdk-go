package awsquicksight


// Determines the row alternate color options.
//
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
type CfnTemplate_RowAlternateColorOptionsProperty struct {
	// Determines the list of row alternate colors.
	RowAlternateColors *[]*string `field:"optional" json:"rowAlternateColors" yaml:"rowAlternateColors"`
	// Determines the widget status.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

