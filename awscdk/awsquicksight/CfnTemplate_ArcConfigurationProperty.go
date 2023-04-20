package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   arcConfigurationProperty := &ArcConfigurationProperty{
//   	ArcAngle: jsii.Number(123),
//   	ArcThickness: jsii.String("arcThickness"),
//   }
//
type CfnTemplate_ArcConfigurationProperty struct {
	// `CfnTemplate.ArcConfigurationProperty.ArcAngle`.
	ArcAngle *float64 `field:"optional" json:"arcAngle" yaml:"arcAngle"`
	// `CfnTemplate.ArcConfigurationProperty.ArcThickness`.
	ArcThickness *string `field:"optional" json:"arcThickness" yaml:"arcThickness"`
}

