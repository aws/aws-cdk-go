package previewawsosismixins


// Options for attaching a VPC to pipeline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcAttachmentOptionsProperty := &VpcAttachmentOptionsProperty{
//   	AttachToVpc: jsii.Boolean(false),
//   	CidrBlock: jsii.String("cidrBlock"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-vpcattachmentoptions.html
//
type CfnPipelinePropsMixin_VpcAttachmentOptionsProperty struct {
	// Whether a VPC is attached to the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-vpcattachmentoptions.html#cfn-osis-pipeline-vpcattachmentoptions-attachtovpc
	//
	AttachToVpc interface{} `field:"optional" json:"attachToVpc" yaml:"attachToVpc"`
	// The CIDR block to be reserved for OpenSearch Ingestion to create elastic network interfaces (ENIs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-osis-pipeline-vpcattachmentoptions.html#cfn-osis-pipeline-vpcattachmentoptions-cidrblock
	//
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
}

