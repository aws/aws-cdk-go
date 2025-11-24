package mixinsawsomics


// Contains information about the source reference in a code repository, such as a branch, tag, or commit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceReferenceProperty := &SourceReferenceProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflow-sourcereference.html
//
type CfnWorkflowPropsMixin_SourceReferenceProperty struct {
	// The type of source reference, such as branch, tag, or commit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflow-sourcereference.html#cfn-omics-workflow-sourcereference-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value of the source reference, such as the branch name, tag name, or commit ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-workflow-sourcereference.html#cfn-omics-workflow-sourcereference-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

