package awsssmcontacts


// Information about when an on-call shift begins and ends.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coverageTimeProperty := &CoverageTimeProperty{
//   	EndTime: jsii.String("endTime"),
//   	StartTime: jsii.String("startTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-coveragetime.html
//
type CfnRotation_CoverageTimeProperty struct {
	// Information about when an on-call rotation shift ends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-coveragetime.html#cfn-ssmcontacts-rotation-coveragetime-endtime
	//
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Information about when an on-call rotation shift begins.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-rotation-coveragetime.html#cfn-ssmcontacts-rotation-coveragetime-starttime
	//
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

