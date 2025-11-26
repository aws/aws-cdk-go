package previewawscleanroomsmixins


// Specifies the name of the column that contains the unique identifier of your users, whose privacy you want to protect.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   differentialPrivacyColumnProperty := &DifferentialPrivacyColumnProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-differentialprivacycolumn.html
//
type CfnConfiguredTablePropsMixin_DifferentialPrivacyColumnProperty struct {
	// The name of the column, such as user_id, that contains the unique identifier of your users, whose privacy you want to protect.
	//
	// If you want to turn on differential privacy for two or more tables in a collaboration, you must configure the same column as the user identifier column in both analysis rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-differentialprivacycolumn.html#cfn-cleanrooms-configuredtable-differentialprivacycolumn-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

