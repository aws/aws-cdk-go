package previewawsiotmixins


// The properties of a billing group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   billingGroupPropertiesProperty := &BillingGroupPropertiesProperty{
//   	BillingGroupDescription: jsii.String("billingGroupDescription"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-billinggroup-billinggroupproperties.html
//
type CfnBillingGroupPropsMixin_BillingGroupPropertiesProperty struct {
	// The description of the billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-billinggroup-billinggroupproperties.html#cfn-iot-billinggroup-billinggroupproperties-billinggroupdescription
	//
	BillingGroupDescription *string `field:"optional" json:"billingGroupDescription" yaml:"billingGroupDescription"`
}

