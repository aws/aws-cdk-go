package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputColumnThresholdProperty := &OutputColumnThresholdProperty{
//   	MinimumIdentityCount: jsii.Number(123),
//   	OutputColumnName: jsii.String("outputColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-outputcolumnthreshold.html
//
type CfnIntermediateTable_OutputColumnThresholdProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-outputcolumnthreshold.html#cfn-cleanrooms-intermediatetable-outputcolumnthreshold-minimumidentitycount
	//
	MinimumIdentityCount *float64 `field:"required" json:"minimumIdentityCount" yaml:"minimumIdentityCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-intermediatetable-outputcolumnthreshold.html#cfn-cleanrooms-intermediatetable-outputcolumnthreshold-outputcolumnname
	//
	OutputColumnName *string `field:"required" json:"outputColumnName" yaml:"outputColumnName"`
}

