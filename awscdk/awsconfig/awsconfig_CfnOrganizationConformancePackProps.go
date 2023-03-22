package awsconfig


// Properties for defining a `CfnOrganizationConformancePack`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationConformancePackProps := &cfnOrganizationConformancePackProps{
//   	organizationConformancePackName: jsii.String("organizationConformancePackName"),
//
//   	// the properties below are optional
//   	conformancePackInputParameters: []interface{}{
//   		&conformancePackInputParameterProperty{
//   			parameterName: jsii.String("parameterName"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	deliveryS3Bucket: jsii.String("deliveryS3Bucket"),
//   	deliveryS3KeyPrefix: jsii.String("deliveryS3KeyPrefix"),
//   	excludedAccounts: []*string{
//   		jsii.String("excludedAccounts"),
//   	},
//   	templateBody: jsii.String("templateBody"),
//   	templateS3Uri: jsii.String("templateS3Uri"),
//   }
//
type CfnOrganizationConformancePackProps struct {
	// The name you assign to an organization conformance pack.
	OrganizationConformancePackName *string `field:"required" json:"organizationConformancePackName" yaml:"organizationConformancePackName"`
	// A list of `ConformancePackInputParameter` objects.
	ConformancePackInputParameters interface{} `field:"optional" json:"conformancePackInputParameters" yaml:"conformancePackInputParameters"`
	// The name of the Amazon S3 bucket where AWS Config stores conformance pack templates.
	//
	// > This field is optional.
	DeliveryS3Bucket *string `field:"optional" json:"deliveryS3Bucket" yaml:"deliveryS3Bucket"`
	// Any folder structure you want to add to an Amazon S3 bucket.
	//
	// > This field is optional.
	DeliveryS3KeyPrefix *string `field:"optional" json:"deliveryS3KeyPrefix" yaml:"deliveryS3KeyPrefix"`
	// A comma-separated list of accounts excluded from organization conformance pack.
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// A string containing full conformance pack template body.
	//
	// Structure containing the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Location of file containing the template body.
	//
	// The uri must point to the conformance pack template (max size: 300 KB).
	TemplateS3Uri *string `field:"optional" json:"templateS3Uri" yaml:"templateS3Uri"`
}

