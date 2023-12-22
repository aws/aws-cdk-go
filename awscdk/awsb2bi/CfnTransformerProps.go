package awsb2bi

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransformer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransformerProps := &CfnTransformerProps{
//   	EdiType: &EdiTypeProperty{
//   		X12Details: &X12DetailsProperty{
//   			TransactionSet: jsii.String("transactionSet"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	FileFormat: jsii.String("fileFormat"),
//   	MappingTemplate: jsii.String("mappingTemplate"),
//   	Name: jsii.String("name"),
//   	Status: jsii.String("status"),
//
//   	// the properties below are optional
//   	ModifiedAt: jsii.String("modifiedAt"),
//   	SampleDocument: jsii.String("sampleDocument"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html
//
type CfnTransformerProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-editype
	//
	EdiType interface{} `field:"required" json:"ediType" yaml:"ediType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-fileformat
	//
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-mappingtemplate
	//
	MappingTemplate *string `field:"required" json:"mappingTemplate" yaml:"mappingTemplate"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-modifiedat
	//
	ModifiedAt *string `field:"optional" json:"modifiedAt" yaml:"modifiedAt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-sampledocument
	//
	SampleDocument *string `field:"optional" json:"sampleDocument" yaml:"sampleDocument"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-b2bi-transformer.html#cfn-b2bi-transformer-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

