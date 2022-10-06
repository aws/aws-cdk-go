package awsmacie


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   criteriaProperty := &criteriaProperty{
//   	regex: jsii.String("regex"),
//   	s3WordsList: &s3WordsListProperty{
//   		bucketName: jsii.String("bucketName"),
//   		objectKey: jsii.String("objectKey"),
//   	},
//   }
//
type CfnAllowList_CriteriaProperty struct {
	// `CfnAllowList.CriteriaProperty.Regex`.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// `CfnAllowList.CriteriaProperty.S3WordsList`.
	S3WordsList interface{} `field:"optional" json:"s3WordsList" yaml:"s3WordsList"`
}

