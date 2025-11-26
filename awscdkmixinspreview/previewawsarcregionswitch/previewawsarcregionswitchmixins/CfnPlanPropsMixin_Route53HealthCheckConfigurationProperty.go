package previewawsarcregionswitchmixins


// The Amazon Route 53 health check configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   route53HealthCheckConfigurationProperty := &Route53HealthCheckConfigurationProperty{
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   	RecordName: jsii.String("recordName"),
//   	RecordSets: []interface{}{
//   		&Route53ResourceRecordSetProperty{
//   			RecordSetIdentifier: jsii.String("recordSetIdentifier"),
//   			Region: jsii.String("region"),
//   		},
//   	},
//   	TimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthcheckconfiguration.html
//
type CfnPlanPropsMixin_Route53HealthCheckConfigurationProperty struct {
	// The cross account role for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthcheckconfiguration.html#cfn-arcregionswitch-plan-route53healthcheckconfiguration-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// The external ID (secret key) for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthcheckconfiguration.html#cfn-arcregionswitch-plan-route53healthcheckconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The Amazon Route 53 health check configuration hosted zone ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthcheckconfiguration.html#cfn-arcregionswitch-plan-route53healthcheckconfiguration-hostedzoneid
	//
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The Amazon Route 53 health check configuration record name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthcheckconfiguration.html#cfn-arcregionswitch-plan-route53healthcheckconfiguration-recordname
	//
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
	// The Amazon Route 53 health check configuration record sets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthcheckconfiguration.html#cfn-arcregionswitch-plan-route53healthcheckconfiguration-recordsets
	//
	RecordSets interface{} `field:"optional" json:"recordSets" yaml:"recordSets"`
	// The Amazon Route 53 health check configuration time out (in minutes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-route53healthcheckconfiguration.html#cfn-arcregionswitch-plan-route53healthcheckconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

