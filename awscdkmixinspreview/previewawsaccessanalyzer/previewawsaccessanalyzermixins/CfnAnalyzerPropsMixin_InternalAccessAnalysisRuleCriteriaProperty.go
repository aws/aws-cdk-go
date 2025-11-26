package previewawsaccessanalyzermixins


// The criteria for an analysis rule for an internal access analyzer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   internalAccessAnalysisRuleCriteriaProperty := &InternalAccessAnalysisRuleCriteriaProperty{
//   	AccountIds: []*string{
//   		jsii.String("accountIds"),
//   	},
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-internalaccessanalysisrulecriteria.html
//
type CfnAnalyzerPropsMixin_InternalAccessAnalysisRuleCriteriaProperty struct {
	// A list of AWS account IDs to apply to the internal access analysis rule criteria.
	//
	// Account IDs can only be applied to the analysis rule criteria for organization-level analyzers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-internalaccessanalysisrulecriteria.html#cfn-accessanalyzer-analyzer-internalaccessanalysisrulecriteria-accountids
	//
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// A list of resource ARNs to apply to the internal access analysis rule criteria.
	//
	// The analyzer will only generate findings for resources that match these ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-internalaccessanalysisrulecriteria.html#cfn-accessanalyzer-analyzer-internalaccessanalysisrulecriteria-resourcearns
	//
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// A list of resource types to apply to the internal access analysis rule criteria.
	//
	// The analyzer will only generate findings for resources of these types. These resource types are currently supported for internal access analyzers:
	//
	// - `AWS::S3::Bucket`
	// - `AWS::RDS::DBSnapshot`
	// - `AWS::RDS::DBClusterSnapshot`
	// - `AWS::S3Express::DirectoryBucket`
	// - `AWS::DynamoDB::Table`
	// - `AWS::DynamoDB::Stream`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-analyzer-internalaccessanalysisrulecriteria.html#cfn-accessanalyzer-analyzer-internalaccessanalysisrulecriteria-resourcetypes
	//
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

