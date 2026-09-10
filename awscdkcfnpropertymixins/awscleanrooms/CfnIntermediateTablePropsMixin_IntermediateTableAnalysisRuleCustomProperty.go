package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   intermediateTableAnalysisRuleCustomProperty := &IntermediateTableAnalysisRuleCustomProperty{
//   	AdditionalAnalyses: jsii.String("additionalAnalyses"),
//   	AggregationThresholds: []interface{}{
//   		&AggregationThresholdProperty{
//   			AllowedAggregateExpressionType: jsii.String("allowedAggregateExpressionType"),
//   			IdentityColumns: []*string{
//   				jsii.String("identityColumns"),
//   			},
//   			MinimumIdentityCount: jsii.Number(123),
//   			OutputColumnThresholds: []interface{}{
//   				&OutputColumnThresholdProperty{
//   					MinimumIdentityCount: jsii.Number(123),
//   					OutputColumnName: jsii.String("outputColumnName"),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	AllowedAnalyses: []*string{
//   		jsii.String("allowedAnalyses"),
//   	},
//   	AllowedAnalysisProviders: []*string{
//   		jsii.String("allowedAnalysisProviders"),
//   	},
//   	AllowedResultReceivers: []*string{
//   		jsii.String("allowedResultReceivers"),
//   	},
//   	ComparisonControls: &ComparisonControlsProperty{
//   		AllowedColumnComparisonColumns: []*string{
//   			jsii.String("allowedColumnComparisonColumns"),
//   		},
//   		AllowedLiteralComparisonColumns: []*string{
//   			jsii.String("allowedLiteralComparisonColumns"),
//   		},
//   	},
//   	DifferentialPrivacy: &DifferentialPrivacyProperty{
//   		Columns: []interface{}{
//   			&DifferentialPrivacyColumnProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   	},
//   	DisallowedOutputColumns: []*string{
//   		jsii.String("disallowedOutputColumns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html
//
type CfnIntermediateTablePropsMixin_IntermediateTableAnalysisRuleCustomProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-additionalanalyses
	//
	AdditionalAnalyses *string `field:"optional" json:"additionalAnalyses" yaml:"additionalAnalyses"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-aggregationthresholds
	//
	AggregationThresholds interface{} `field:"optional" json:"aggregationThresholds" yaml:"aggregationThresholds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-allowedanalyses
	//
	AllowedAnalyses *[]*string `field:"optional" json:"allowedAnalyses" yaml:"allowedAnalyses"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-allowedanalysisproviders
	//
	AllowedAnalysisProviders *[]*string `field:"optional" json:"allowedAnalysisProviders" yaml:"allowedAnalysisProviders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-allowedresultreceivers
	//
	AllowedResultReceivers *[]*string `field:"optional" json:"allowedResultReceivers" yaml:"allowedResultReceivers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-comparisoncontrols
	//
	ComparisonControls interface{} `field:"optional" json:"comparisonControls" yaml:"comparisonControls"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-differentialprivacy
	//
	DifferentialPrivacy interface{} `field:"optional" json:"differentialPrivacy" yaml:"differentialPrivacy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom.html#cfn-cleanrooms-intermediatetable-intermediatetableanalysisrulecustom-disallowedoutputcolumns
	//
	DisallowedOutputColumns *[]*string `field:"optional" json:"disallowedOutputColumns" yaml:"disallowedOutputColumns"`
}

