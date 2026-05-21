package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnSemanticTypeProperty := &ColumnSemanticTypeProperty{
//   	GeographicalRole: jsii.String("geographicalRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnsemantictype.html
//
type CfnDataSet_ColumnSemanticTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-columnsemantictype.html#cfn-quicksight-dataset-columnsemantictype-geographicalrole
	//
	GeographicalRole *string `field:"optional" json:"geographicalRole" yaml:"geographicalRole"`
}

