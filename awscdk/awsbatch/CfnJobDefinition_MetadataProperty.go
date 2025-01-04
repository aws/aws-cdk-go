package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var labels interface{}
//
//   metadataProperty := &MetadataProperty{
//   	Labels: labels,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-metadata.html
//
type CfnJobDefinition_MetadataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-metadata.html#cfn-batch-jobdefinition-metadata-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
}

