package awssagemaker


// The number of cluster nodes that the extension applies to, and how many of them run the requested extension version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   extensionNodeSummaryProperty := &ExtensionNodeSummaryProperty{
//   	ActiveCount: jsii.Number(123),
//   	TotalCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-extension-extensionnodesummary.html
//
type CfnExtension_ExtensionNodeSummaryProperty struct {
	// The number of cluster nodes that run the requested extension version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-extension-extensionnodesummary.html#cfn-sagemaker-extension-extensionnodesummary-activecount
	//
	ActiveCount *float64 `field:"required" json:"activeCount" yaml:"activeCount"`
	// The number of cluster nodes that the extension applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-extension-extensionnodesummary.html#cfn-sagemaker-extension-extensionnodesummary-totalcount
	//
	TotalCount *float64 `field:"required" json:"totalCount" yaml:"totalCount"`
}

