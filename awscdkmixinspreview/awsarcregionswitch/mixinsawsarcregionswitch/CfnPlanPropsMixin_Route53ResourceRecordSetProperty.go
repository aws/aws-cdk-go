package mixinsawsarcregionswitch


// The Amazon Route 53 record set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53ResourceRecordSetProperty := &Route53ResourceRecordSetProperty{
//   	RecordSetIdentifier: jsii.String("recordSetIdentifier"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53resourcerecordset.html
//
type CfnPlanPropsMixin_Route53ResourceRecordSetProperty struct {
	// The Amazon Route 53 record set identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53resourcerecordset.html#cfn-arcregionswitch-plan-route53resourcerecordset-recordsetidentifier
	//
	RecordSetIdentifier *string `field:"optional" json:"recordSetIdentifier" yaml:"recordSetIdentifier"`
	// The Amazon Route 53 record set Region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53resourcerecordset.html#cfn-arcregionswitch-plan-route53resourcerecordset-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

