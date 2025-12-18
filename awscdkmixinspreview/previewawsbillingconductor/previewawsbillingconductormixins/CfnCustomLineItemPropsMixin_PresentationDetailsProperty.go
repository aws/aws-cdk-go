package previewawsbillingconductormixins


// An object that defines how custom line item charges are presented in the bill, containing specifications for service presentation.
//
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
	// The service under which the custom line item charges will be presented.
	//
	// Must be a string between 1 and 128 characters matching the pattern `^[a-zA-Z0-9]+$` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-billingconductor-customlineitem-presentationdetails.html#cfn-billingconductor-customlineitem-presentationdetails-service
	//
	Service *string `field:"optional" json:"service" yaml:"service"`
}

