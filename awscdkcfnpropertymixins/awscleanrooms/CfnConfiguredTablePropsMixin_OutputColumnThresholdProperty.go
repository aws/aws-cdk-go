package awscleanrooms


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   outputColumnThresholdProperty := &OutputColumnThresholdProperty{
//   	MinimumIdentityCount: jsii.Number(123),
//   	OutputColumnName: jsii.String("outputColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-outputcolumnthreshold.html
//
type CfnConfiguredTablePropsMixin_OutputColumnThresholdProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-outputcolumnthreshold.html#cfn-cleanrooms-configuredtable-outputcolumnthreshold-minimumidentitycount
	//
	MinimumIdentityCount *float64 `field:"optional" json:"minimumIdentityCount" yaml:"minimumIdentityCount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cleanrooms-configuredtable-outputcolumnthreshold.html#cfn-cleanrooms-configuredtable-outputcolumnthreshold-outputcolumnname
	//
	OutputColumnName *string `field:"optional" json:"outputColumnName" yaml:"outputColumnName"`
}

