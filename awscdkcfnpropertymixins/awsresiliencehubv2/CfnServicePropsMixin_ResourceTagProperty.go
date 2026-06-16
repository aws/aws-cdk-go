package awsresiliencehubv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   resourceTagProperty := &ResourceTagProperty{
//   	Key: jsii.String("key"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-resourcetag.html
//
type CfnServicePropsMixin_ResourceTagProperty struct {
	// Tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-resourcetag.html#cfn-resiliencehubv2-service-resourcetag-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Tag values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehubv2-service-resourcetag.html#cfn-resiliencehubv2-service-resourcetag-values
	//
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

