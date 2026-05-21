package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   uploadedDocumentMetadataProperty := &UploadedDocumentMetadataProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadeddocumentmetadata.html
//
type CfnDataSet_UploadedDocumentMetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-uploadeddocumentmetadata.html#cfn-quicksight-dataset-uploadeddocumentmetadata-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

