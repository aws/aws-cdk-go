package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnIntermediateTablePropsMixin_DifferentialPrivacyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-differentialprivacy.html#cfn-cleanrooms-intermediatetable-differentialprivacy-columns
	//
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
}

