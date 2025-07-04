package awsaccessanalyzer


// Specifies the configuration of an internal access analyzer for an AWS organization or account.
//
// This configuration determines how the analyzer evaluates internal access within your AWS environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   internalAccessConfigurationProperty := &InternalAccessConfigurationProperty{
//   	InternalAccessAnalysisRule: &InternalAccessAnalysisRuleProperty{
//   		Inclusions: []interface{}{
//   			&InternalAccessAnalysisRuleCriteriaProperty{
//   				AccountIds: []*string{
//   					jsii.String("accountIds"),
//   				},
//   				ResourceArns: []*string{
//   					jsii.String("resourceArns"),
//   				},
//   				ResourceTypes: []*string{
//   					jsii.String("resourceTypes"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-internalaccessconfiguration.html
//
type CfnAnalyzer_InternalAccessConfigurationProperty struct {
	// Contains information about analysis rules for the internal access analyzer.
	//
	// These rules determine which resources and access patterns will be analyzed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-internalaccessconfiguration.html#cfn-accessanalyzer-analyzer-internalaccessconfiguration-internalaccessanalysisrule
	//
	InternalAccessAnalysisRule interface{} `field:"optional" json:"internalAccessAnalysisRule" yaml:"internalAccessAnalysisRule"`
}

