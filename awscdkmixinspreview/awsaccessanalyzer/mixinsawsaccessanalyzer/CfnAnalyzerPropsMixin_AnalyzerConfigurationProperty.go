package mixinsawsaccessanalyzer


// Contains information about the configuration of an analyzer for an AWS organization or account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   analyzerConfigurationProperty := &AnalyzerConfigurationProperty{
//   	InternalAccessConfiguration: &InternalAccessConfigurationProperty{
//   		InternalAccessAnalysisRule: &InternalAccessAnalysisRuleProperty{
//   			Inclusions: []interface{}{
//   				&InternalAccessAnalysisRuleCriteriaProperty{
//   					AccountIds: []*string{
//   						jsii.String("accountIds"),
//   					},
//   					ResourceArns: []*string{
//   						jsii.String("resourceArns"),
//   					},
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	UnusedAccessConfiguration: &UnusedAccessConfigurationProperty{
//   		AnalysisRule: &AnalysisRuleProperty{
//   			Exclusions: []interface{}{
//   				&AnalysisRuleCriteriaProperty{
//   					AccountIds: []*string{
//   						jsii.String("accountIds"),
//   					},
//   					ResourceTags: []interface{}{
//   						[]interface{}{
//   							&CfnTag{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   		UnusedAccessAge: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analyzerconfiguration.html
//
type CfnAnalyzerPropsMixin_AnalyzerConfigurationProperty struct {
	// Specifies the configuration of an internal access analyzer for an AWS organization or account.
	//
	// This configuration determines how the analyzer evaluates access within your AWS environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analyzerconfiguration.html#cfn-accessanalyzer-analyzer-analyzerconfiguration-internalaccessconfiguration
	//
	InternalAccessConfiguration interface{} `field:"optional" json:"internalAccessConfiguration" yaml:"internalAccessConfiguration"`
	// Specifies the configuration of an unused access analyzer for an AWS organization or account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-analyzerconfiguration.html#cfn-accessanalyzer-analyzer-analyzerconfiguration-unusedaccessconfiguration
	//
	UnusedAccessConfiguration interface{} `field:"optional" json:"unusedAccessConfiguration" yaml:"unusedAccessConfiguration"`
}

