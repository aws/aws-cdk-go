package awss3


// This resource contains `DaysGreaterThan` and `DaysLessThan` to define the object age range (minimum and maximum number of days).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchObjectAgeProperty := &MatchObjectAgeProperty{
//   	DaysGreaterThan: jsii.Number(123),
//   	DaysLessThan: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectage.html
//
type CfnStorageLensGroup_MatchObjectAgeProperty struct {
	// This property indicates the minimum object age in days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectage.html#cfn-s3-storagelensgroup-matchobjectage-daysgreaterthan
	//
	DaysGreaterThan *float64 `field:"optional" json:"daysGreaterThan" yaml:"daysGreaterThan"`
	// This property indicates the maximum object age in days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectage.html#cfn-s3-storagelensgroup-matchobjectage-dayslessthan
	//
	DaysLessThan *float64 `field:"optional" json:"daysLessThan" yaml:"daysLessThan"`
}

