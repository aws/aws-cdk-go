package awsconfig


// Properties for defining a `CfnConformancePack`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var templateSsmDocumentDetails interface{}
//
//   cfnConformancePackProps := &CfnConformancePackProps{
//   	ConformancePackName: jsii.String("conformancePackName"),
//
//   	// the properties below are optional
//   	ConformancePackInputParameters: []interface{}{
//   		&ConformancePackInputParameterProperty{
//   			ParameterName: jsii.String("parameterName"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	DeliveryS3Bucket: jsii.String("deliveryS3Bucket"),
//   	DeliveryS3KeyPrefix: jsii.String("deliveryS3KeyPrefix"),
//   	TemplateBody: jsii.String("templateBody"),
//   	TemplateS3Uri: jsii.String("templateS3Uri"),
//   	TemplateSsmDocumentDetails: templateSsmDocumentDetails,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html
//
type CfnConformancePackProps struct {
	// Name of the conformance pack you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-conformancepackname
	//
	ConformancePackName *string `field:"required" json:"conformancePackName" yaml:"conformancePackName"`
	// A list of ConformancePackInputParameter objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-conformancepackinputparameters
	//
	ConformancePackInputParameters interface{} `field:"optional" json:"conformancePackInputParameters" yaml:"conformancePackInputParameters"`
	// The name of the Amazon S3 bucket where AWS Config stores conformance pack templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-deliverys3bucket
	//
	DeliveryS3Bucket *string `field:"optional" json:"deliveryS3Bucket" yaml:"deliveryS3Bucket"`
	// The prefix for the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-deliverys3keyprefix
	//
	DeliveryS3KeyPrefix *string `field:"optional" json:"deliveryS3KeyPrefix" yaml:"deliveryS3KeyPrefix"`
	// A string containing full conformance pack template body.
	//
	// Structure containing the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	//
	// > You can only use a YAML template with two resource types: config rule ( `AWS::Config::ConfigRule` ) and a remediation action ( `AWS::Config::RemediationConfiguration` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-templatebody
	//
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Location of file containing the template body (s3://bucketname/prefix).
	//
	// The uri must point to the conformance pack template (max size: 300 KB) that is located in an Amazon S3 bucket.
	//
	// > You must have access to read Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-templates3uri
	//
	TemplateS3Uri *string `field:"optional" json:"templateS3Uri" yaml:"templateS3Uri"`
	// An object that contains the name or Amazon Resource Name (ARN) of the AWS Systems Manager document (SSM document) and the version of the SSM document that is used to create a conformance pack.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-conformancepack.html#cfn-config-conformancepack-templatessmdocumentdetails
	//
	TemplateSsmDocumentDetails interface{} `field:"optional" json:"templateSsmDocumentDetails" yaml:"templateSsmDocumentDetails"`
}

