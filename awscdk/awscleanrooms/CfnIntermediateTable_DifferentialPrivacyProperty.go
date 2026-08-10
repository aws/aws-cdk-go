package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   differentialPrivacyProperty := &DifferentialPrivacyProperty{
//   	Columns: []interface{}{
//   		&DifferentialPrivacyColumnProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-differentialprivacy.html
//
type CfnIntermediateTable_DifferentialPrivacyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-differentialprivacy.html#cfn-cleanrooms-intermediatetable-differentialprivacy-columns
	//
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
}

