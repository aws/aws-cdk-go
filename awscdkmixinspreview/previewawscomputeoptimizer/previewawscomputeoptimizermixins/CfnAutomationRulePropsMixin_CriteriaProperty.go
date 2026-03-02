package previewawscomputeoptimizermixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   criteriaProperty := &CriteriaProperty{
//   	EbsVolumeSizeInGib: []interface{}{
//   		&IntegerCriteriaConditionProperty{
//   			Comparison: jsii.String("comparison"),
//   			Values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	EbsVolumeType: []interface{}{
//   		&StringCriteriaConditionProperty{
//   			Comparison: jsii.String("comparison"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	EstimatedMonthlySavings: []interface{}{
//   		&DoubleCriteriaConditionProperty{
//   			Comparison: jsii.String("comparison"),
//   			Values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	LookBackPeriodInDays: []interface{}{
//   		&IntegerCriteriaConditionProperty{
//   			Comparison: jsii.String("comparison"),
//   			Values: []interface{}{
//   				jsii.Number(123),
//   			},
//   		},
//   	},
//   	Region: []interface{}{
//   		&StringCriteriaConditionProperty{
//   			Comparison: jsii.String("comparison"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	ResourceArn: []interface{}{
//   		&StringCriteriaConditionProperty{
//   			Comparison: jsii.String("comparison"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	ResourceTag: []interface{}{
//   		&ResourceTagsCriteriaConditionProperty{
//   			Comparison: jsii.String("comparison"),
//   			Key: jsii.String("key"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   	RestartNeeded: []interface{}{
//   		&StringCriteriaConditionProperty{
//   			Comparison: jsii.String("comparison"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html
//
type CfnAutomationRulePropsMixin_CriteriaProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-ebsvolumesizeingib
	//
	EbsVolumeSizeInGib interface{} `field:"optional" json:"ebsVolumeSizeInGib" yaml:"ebsVolumeSizeInGib"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-ebsvolumetype
	//
	EbsVolumeType interface{} `field:"optional" json:"ebsVolumeType" yaml:"ebsVolumeType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-estimatedmonthlysavings
	//
	EstimatedMonthlySavings interface{} `field:"optional" json:"estimatedMonthlySavings" yaml:"estimatedMonthlySavings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-lookbackperiodindays
	//
	LookBackPeriodInDays interface{} `field:"optional" json:"lookBackPeriodInDays" yaml:"lookBackPeriodInDays"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-region
	//
	Region interface{} `field:"optional" json:"region" yaml:"region"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-resourcearn
	//
	ResourceArn interface{} `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-resourcetag
	//
	ResourceTag interface{} `field:"optional" json:"resourceTag" yaml:"resourceTag"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-computeoptimizer-automationrule-criteria.html#cfn-computeoptimizer-automationrule-criteria-restartneeded
	//
	RestartNeeded interface{} `field:"optional" json:"restartNeeded" yaml:"restartNeeded"`
}

