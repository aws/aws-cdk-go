package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceSpecificationProperty := &InferenceSpecificationProperty{
//   	Containers: []interface{}{
//   		&ContainerProperty{
//   			Image: jsii.String("image"),
//
//   			// the properties below are optional
//   			ModelDataUrl: jsii.String("modelDataUrl"),
//   			NearestModelName: jsii.String("nearestModelName"),
//   		},
//   	},
//   }
//
type CfnModelCard_InferenceSpecificationProperty struct {
	// `CfnModelCard.InferenceSpecificationProperty.Containers`.
	Containers interface{} `field:"required" json:"containers" yaml:"containers"`
}

