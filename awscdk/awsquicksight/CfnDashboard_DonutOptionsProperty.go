package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   donutOptionsProperty := &DonutOptionsProperty{
//   	ArcOptions: &ArcOptionsProperty{
//   		ArcThickness: jsii.String("arcThickness"),
//   	},
//   	DonutCenterOptions: &DonutCenterOptionsProperty{
//   		LabelVisibility: jsii.String("labelVisibility"),
//   	},
//   }
//
type CfnDashboard_DonutOptionsProperty struct {
	// `CfnDashboard.DonutOptionsProperty.ArcOptions`.
	ArcOptions interface{} `field:"optional" json:"arcOptions" yaml:"arcOptions"`
	// `CfnDashboard.DonutOptionsProperty.DonutCenterOptions`.
	DonutCenterOptions interface{} `field:"optional" json:"donutCenterOptions" yaml:"donutCenterOptions"`
}

