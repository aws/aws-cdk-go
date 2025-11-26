package previewawsarcregionswitchmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53HealthChecksProperty := &Route53HealthChecksProperty{
//   	HealthCheckIds: []*string{
//   		jsii.String("healthCheckIds"),
//   	},
//   	HostedZoneIds: []*string{
//   		jsii.String("hostedZoneIds"),
//   	},
//   	RecordNames: []*string{
//   		jsii.String("recordNames"),
//   	},
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthchecks.html
//
type CfnPlanPropsMixin_Route53HealthChecksProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthchecks.html#cfn-arcregionswitch-plan-route53healthchecks-healthcheckids
	//
	HealthCheckIds *[]*string `field:"optional" json:"healthCheckIds" yaml:"healthCheckIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthchecks.html#cfn-arcregionswitch-plan-route53healthchecks-hostedzoneids
	//
	HostedZoneIds *[]*string `field:"optional" json:"hostedZoneIds" yaml:"hostedZoneIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthchecks.html#cfn-arcregionswitch-plan-route53healthchecks-recordnames
	//
	RecordNames *[]*string `field:"optional" json:"recordNames" yaml:"recordNames"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthchecks.html#cfn-arcregionswitch-plan-route53healthchecks-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

