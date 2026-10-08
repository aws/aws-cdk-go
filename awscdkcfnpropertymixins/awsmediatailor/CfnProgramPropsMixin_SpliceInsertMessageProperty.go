package awsmediatailor


// Splice insert message configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   spliceInsertMessageProperty := &SpliceInsertMessageProperty{
//   	AvailNum: jsii.Number(123),
//   	AvailsExpected: jsii.Number(123),
//   	SpliceEventId: jsii.Number(123),
//   	UniqueProgramId: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-spliceinsertmessage.html
//
type CfnProgramPropsMixin_SpliceInsertMessageProperty struct {
	// This is written to splice_insert.avail_num.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-spliceinsertmessage.html#cfn-mediatailor-program-spliceinsertmessage-availnum
	//
	AvailNum *float64 `field:"optional" json:"availNum" yaml:"availNum"`
	// This is written to splice_insert.avails_expected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-spliceinsertmessage.html#cfn-mediatailor-program-spliceinsertmessage-availsexpected
	//
	AvailsExpected *float64 `field:"optional" json:"availsExpected" yaml:"availsExpected"`
	// This is written to splice_insert.splice_event_id.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-spliceinsertmessage.html#cfn-mediatailor-program-spliceinsertmessage-spliceeventid
	//
	SpliceEventId *float64 `field:"optional" json:"spliceEventId" yaml:"spliceEventId"`
	// This is written to splice_insert.unique_program_id.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-program-spliceinsertmessage.html#cfn-mediatailor-program-spliceinsertmessage-uniqueprogramid
	//
	UniqueProgramId *float64 `field:"optional" json:"uniqueProgramId" yaml:"uniqueProgramId"`
}

