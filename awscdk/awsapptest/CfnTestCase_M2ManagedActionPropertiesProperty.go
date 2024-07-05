package awsapptest


// Specifies the AWS Mainframe Modernization managed action properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   m2ManagedActionPropertiesProperty := &M2ManagedActionPropertiesProperty{
//   	ForceStop: jsii.Boolean(false),
//   	ImportDataSetLocation: jsii.String("importDataSetLocation"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedactionproperties.html
//
type CfnTestCase_M2ManagedActionPropertiesProperty struct {
	// Force stops the AWS Mainframe Modernization managed action properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedactionproperties.html#cfn-apptest-testcase-m2managedactionproperties-forcestop
	//
	ForceStop interface{} `field:"optional" json:"forceStop" yaml:"forceStop"`
	// The import data set location of the AWS Mainframe Modernization managed action properties.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedactionproperties.html#cfn-apptest-testcase-m2managedactionproperties-importdatasetlocation
	//
	ImportDataSetLocation *string `field:"optional" json:"importDataSetLocation" yaml:"importDataSetLocation"`
}

