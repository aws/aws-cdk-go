package awss3files


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   importDataRuleProperty := &ImportDataRuleProperty{
//   	Prefix: jsii.String("prefix"),
//   	SizeLessThan: jsii.Number(123),
//   	Trigger: jsii.String("trigger"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-importdatarule.html
//
type CfnFileSystemPropsMixin_ImportDataRuleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-importdatarule.html#cfn-s3files-filesystem-importdatarule-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-importdatarule.html#cfn-s3files-filesystem-importdatarule-sizelessthan
	//
	SizeLessThan *float64 `field:"optional" json:"sizeLessThan" yaml:"sizeLessThan"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-filesystem-importdatarule.html#cfn-s3files-filesystem-importdatarule-trigger
	//
	Trigger *string `field:"optional" json:"trigger" yaml:"trigger"`
}

