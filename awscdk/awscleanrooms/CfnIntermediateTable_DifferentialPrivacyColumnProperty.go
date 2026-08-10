package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   differentialPrivacyColumnProperty := &DifferentialPrivacyColumnProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-differentialprivacycolumn.html
//
type CfnIntermediateTable_DifferentialPrivacyColumnProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-differentialprivacycolumn.html#cfn-cleanrooms-intermediatetable-differentialprivacycolumn-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

