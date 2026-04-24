package awsses


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   trustStoreProperty := &TrustStoreProperty{
//   	CaContent: jsii.String("caContent"),
//   	CrlContent: jsii.String("crlContent"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-truststore.html
//
type CfnMailManagerIngressPointPropsMixin_TrustStoreProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-truststore.html#cfn-ses-mailmanageringresspoint-truststore-cacontent
	//
	CaContent *string `field:"optional" json:"caContent" yaml:"caContent"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-truststore.html#cfn-ses-mailmanageringresspoint-truststore-crlcontent
	//
	CrlContent *string `field:"optional" json:"crlContent" yaml:"crlContent"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanageringresspoint-truststore.html#cfn-ses-mailmanageringresspoint-truststore-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

