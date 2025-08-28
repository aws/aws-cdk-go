package awsssm


// An SSM document required by the current document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentRequiresProperty := &DocumentRequiresProperty{
//   	Name: jsii.String("name"),
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html
//
type CfnDocument_DocumentRequiresProperty struct {
	// The name of the required SSM document.
	//
	// The name can be an Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The document version required by the current document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-document-documentrequires.html#cfn-ssm-document-documentrequires-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

