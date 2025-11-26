package previewawspcaconnectorscepmixins


// Properties for CfnChallengePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChallengeMixinProps := &CfnChallengeMixinProps{
//   	ConnectorArn: jsii.String("connectorArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-challenge.html
//
type CfnChallengeMixinProps struct {
	// The Amazon Resource Name (ARN) of the connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-challenge.html#cfn-pcaconnectorscep-challenge-connectorarn
	//
	ConnectorArn *string `field:"optional" json:"connectorArn" yaml:"connectorArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pcaconnectorscep-challenge.html#cfn-pcaconnectorscep-challenge-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

