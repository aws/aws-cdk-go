package awsbatch


// Metadata about the Kubernetes pod.
//
// For more information, see [Understanding Kubernetes Objects](https://docs.aws.amazon.com/https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/) in the *Kubernetes documentation* .
//
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
	// Key-value pairs used to identify, sort, and organize cube resources.
	//
	// Can contain up to 63 uppercase letters, lowercase letters, numbers, hyphens (-), and underscores (_). Labels can be added or modified at any time. Each resource can have multiple labels, but each key must be unique for a given object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-metadata.html#cfn-batch-jobdefinition-metadata-labels
	//
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
}

