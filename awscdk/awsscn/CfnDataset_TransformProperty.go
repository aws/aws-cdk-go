package awsscn


// The transformation of the partition field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transformProperty := &TransformProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-transform.html
//
type CfnDataset_TransformProperty struct {
	// The type of partitioning transformation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataset-transform.html#cfn-scn-dataset-transform-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

