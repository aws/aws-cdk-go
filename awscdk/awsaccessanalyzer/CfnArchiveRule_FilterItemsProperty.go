package awsaccessanalyzer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterItemsProperty := &FilterItemsProperty{
//   	Contains: []*string{
//   		jsii.String("contains"),
//   	},
//   	Eq: []*string{
//   		jsii.String("eq"),
//   	},
//   	Exists: jsii.Boolean(false),
//   	Neq: []*string{
//   		jsii.String("neq"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-archiverule-filteritems.html
//
type CfnArchiveRule_FilterItemsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-archiverule-filteritems.html#cfn-accessanalyzer-archiverule-filteritems-contains
	//
	Contains *[]*string `field:"optional" json:"contains" yaml:"contains"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-archiverule-filteritems.html#cfn-accessanalyzer-archiverule-filteritems-eq
	//
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-archiverule-filteritems.html#cfn-accessanalyzer-archiverule-filteritems-exists
	//
	Exists interface{} `field:"optional" json:"exists" yaml:"exists"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-accessanalyzer-archiverule-filteritems.html#cfn-accessanalyzer-archiverule-filteritems-neq
	//
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}

