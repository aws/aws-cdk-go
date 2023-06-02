package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerProperty := &ContainerProperty{
//   	Image: jsii.String("image"),
//
//   	// the properties below are optional
//   	ModelDataUrl: jsii.String("modelDataUrl"),
//   	NearestModelName: jsii.String("nearestModelName"),
//   }
//
type CfnModelCard_ContainerProperty struct {
	// `CfnModelCard.ContainerProperty.Image`.
	Image *string `field:"required" json:"image" yaml:"image"`
	// `CfnModelCard.ContainerProperty.ModelDataUrl`.
	ModelDataUrl *string `field:"optional" json:"modelDataUrl" yaml:"modelDataUrl"`
	// `CfnModelCard.ContainerProperty.NearestModelName`.
	NearestModelName *string `field:"optional" json:"nearestModelName" yaml:"nearestModelName"`
}

