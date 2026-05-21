package awsobservabilityadmin


// Status of a telemetry rule in a specific region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   regionStatusProperty := &RegionStatusProperty{
//   	Region: jsii.String("region"),
//   	RuleArn: jsii.String("ruleArn"),
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-regionstatus.html
//
type CfnOrganizationTelemetryRulePropsMixin_RegionStatusProperty struct {
	// The AWS region code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-regionstatus.html#cfn-observabilityadmin-organizationtelemetryrule-regionstatus-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The ARN of the rule in this region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-regionstatus.html#cfn-observabilityadmin-organizationtelemetryrule-regionstatus-rulearn
	//
	RuleArn *string `field:"optional" json:"ruleArn" yaml:"ruleArn"`
	// The replication status of the rule in this region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationtelemetryrule-regionstatus.html#cfn-observabilityadmin-organizationtelemetryrule-regionstatus-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
}

