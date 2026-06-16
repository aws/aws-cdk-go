package awsbcm


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dateTimeValueProperty := &DateTimeValueProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-datetimevalue.html
//
type CfnDashboardPropsMixin_DateTimeValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-datetimevalue.html#cfn-bcm-dashboard-datetimevalue-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcm-dashboard-datetimevalue.html#cfn-bcm-dashboard-datetimevalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

