package awsdatabrew


// Contains additional resource information needed for specific datasets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataProperty := &metadataProperty{
//   	sourceArn: jsii.String("sourceArn"),
//   }
//
type CfnDataset_MetadataProperty struct {
	// The Amazon Resource Name (ARN) associated with the dataset.
	//
	// Currently, DataBrew only supports ARNs from Amazon AppFlow.
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
}

