package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   analysisTransitGatewayPolicyTableEntryProperty := &AnalysisTransitGatewayPolicyTableEntryProperty{
//   	DestinationCidr: jsii.String("destinationCidr"),
//   	DestinationPortRange: &PortRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   	Protocol: jsii.String("protocol"),
//   	RuleNumber: jsii.Number(123),
//   	SourceCidr: jsii.String("sourceCidr"),
//   	SourcePortRange: &PortRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   	TargetRouteTableId: jsii.String("targetRouteTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry.html
//
type CfnNetworkInsightsAnalysis_AnalysisTransitGatewayPolicyTableEntryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry.html#cfn-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry-destinationcidr
	//
	DestinationCidr *string `field:"optional" json:"destinationCidr" yaml:"destinationCidr"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry.html#cfn-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry-destinationportrange
	//
	DestinationPortRange interface{} `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry.html#cfn-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry.html#cfn-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry-rulenumber
	//
	RuleNumber *float64 `field:"optional" json:"ruleNumber" yaml:"ruleNumber"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry.html#cfn-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry-sourcecidr
	//
	SourceCidr *string `field:"optional" json:"sourceCidr" yaml:"sourceCidr"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry.html#cfn-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry-sourceportrange
	//
	SourcePortRange interface{} `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry.html#cfn-ec2-networkinsightsanalysis-analysistransitgatewaypolicytableentry-targetroutetableid
	//
	TargetRouteTableId *string `field:"optional" json:"targetRouteTableId" yaml:"targetRouteTableId"`
}

