package awscleanrooms


// The analysis method for the configured tables.
//
// The only valid value is currently `DIRECT_QUERY`.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-differentialprivacy.html
//
type CfnConfiguredTable_DifferentialPrivacyProperty struct {
	// The name of the column, such as user_id, that contains the unique identifier of your users, whose privacy you want to protect.
	//
	// If you want to turn on differential privacy for two or more tables in a collaboration, you must configure the same column as the user identifier column in both analysis rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-differentialprivacy.html#cfn-cleanrooms-configuredtable-differentialprivacy-columns
	//
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
}

