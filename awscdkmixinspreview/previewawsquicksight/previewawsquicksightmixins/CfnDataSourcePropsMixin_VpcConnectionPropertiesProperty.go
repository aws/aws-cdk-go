package previewawsquicksightmixins


// VPC connection properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcConnectionPropertiesProperty := &VpcConnectionPropertiesProperty{
//   	VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-vpcconnectionproperties.html
//
type CfnDataSourcePropsMixin_VpcConnectionPropertiesProperty struct {
	// The Amazon Resource Name (ARN) for the VPC connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-vpcconnectionproperties.html#cfn-quicksight-datasource-vpcconnectionproperties-vpcconnectionarn
	//
	VpcConnectionArn *string `field:"optional" json:"vpcConnectionArn" yaml:"vpcConnectionArn"`
}

