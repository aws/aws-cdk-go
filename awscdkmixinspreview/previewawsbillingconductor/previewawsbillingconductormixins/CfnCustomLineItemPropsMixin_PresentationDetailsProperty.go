package previewawsbillingconductormixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   presentationDetailsProperty := &PresentationDetailsProperty{
//   	Service: jsii.String("service"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-presentationdetails.html
//
type CfnCustomLineItemPropsMixin_PresentationDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-presentationdetails.html#cfn-billingconductor-customlineitem-presentationdetails-service
	//
	Service *string `field:"optional" json:"service" yaml:"service"`
}

