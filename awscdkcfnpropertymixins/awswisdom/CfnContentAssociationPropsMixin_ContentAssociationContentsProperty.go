package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   contentAssociationContentsProperty := &ContentAssociationContentsProperty{
//   	AmazonConnectGuideAssociation: &AmazonConnectGuideAssociationDataProperty{
//   		FlowId: jsii.String("flowId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-contentassociation-contentassociationcontents.html
//
type CfnContentAssociationPropsMixin_ContentAssociationContentsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-contentassociation-contentassociationcontents.html#cfn-wisdom-contentassociation-contentassociationcontents-amazonconnectguideassociation
	//
	AmazonConnectGuideAssociation interface{} `field:"optional" json:"amazonConnectGuideAssociation" yaml:"amazonConnectGuideAssociation"`
}

