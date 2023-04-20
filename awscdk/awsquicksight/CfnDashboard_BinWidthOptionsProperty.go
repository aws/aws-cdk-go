package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   binWidthOptionsProperty := &BinWidthOptionsProperty{
//   	BinCountLimit: jsii.Number(123),
//   	Value: jsii.Number(123),
//   }
//
type CfnDashboard_BinWidthOptionsProperty struct {
	// `CfnDashboard.BinWidthOptionsProperty.BinCountLimit`.
	BinCountLimit *float64 `field:"optional" json:"binCountLimit" yaml:"binCountLimit"`
	// `CfnDashboard.BinWidthOptionsProperty.Value`.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

