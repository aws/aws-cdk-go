package mixinsawsomics


// Server-side encryption (SSE) settings for a store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sseConfigProperty := &SseConfigProperty{
//   	KeyArn: jsii.String("keyArn"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-sequencestore-sseconfig.html
//
type CfnSequenceStorePropsMixin_SseConfigProperty struct {
	// An encryption key ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-sequencestore-sseconfig.html#cfn-omics-sequencestore-sseconfig-keyarn
	//
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// The encryption type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-omics-sequencestore-sseconfig.html#cfn-omics-sequencestore-sseconfig-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

