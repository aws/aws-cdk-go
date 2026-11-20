package awsconfig


// The source of a managed conformance pack.
//
// Specifying Source creates a managed conformance pack, whose Config rules are managed by AWS. You can only specify one of Source or the template body or template S3Uri fields.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	SourceIdentifier: jsii.String("sourceIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-source.html
//
type CfnConformancePack_SourceProperty struct {
	// The identifier specifying which managed conformance pack is being deployed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-conformancepack-source.html#cfn-config-conformancepack-source-sourceidentifier
	//
	SourceIdentifier *string `field:"required" json:"sourceIdentifier" yaml:"sourceIdentifier"`
}

