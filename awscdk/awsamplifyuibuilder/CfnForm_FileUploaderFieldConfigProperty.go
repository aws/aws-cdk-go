package awsamplifyuibuilder


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileUploaderFieldConfigProperty := &FileUploaderFieldConfigProperty{
//   	AcceptedFileTypes: []*string{
//   		jsii.String("acceptedFileTypes"),
//   	},
//   	AccessLevel: jsii.String("accessLevel"),
//
//   	// the properties below are optional
//   	IsResumable: jsii.Boolean(false),
//   	MaxFileCount: jsii.Number(123),
//   	MaxSize: jsii.Number(123),
//   	ShowThumbnails: jsii.Boolean(false),
//   }
//
type CfnForm_FileUploaderFieldConfigProperty struct {
	// `CfnForm.FileUploaderFieldConfigProperty.AcceptedFileTypes`.
	AcceptedFileTypes *[]*string `field:"required" json:"acceptedFileTypes" yaml:"acceptedFileTypes"`
	// `CfnForm.FileUploaderFieldConfigProperty.AccessLevel`.
	AccessLevel *string `field:"required" json:"accessLevel" yaml:"accessLevel"`
	// `CfnForm.FileUploaderFieldConfigProperty.IsResumable`.
	IsResumable interface{} `field:"optional" json:"isResumable" yaml:"isResumable"`
	// `CfnForm.FileUploaderFieldConfigProperty.MaxFileCount`.
	MaxFileCount *float64 `field:"optional" json:"maxFileCount" yaml:"maxFileCount"`
	// `CfnForm.FileUploaderFieldConfigProperty.MaxSize`.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// `CfnForm.FileUploaderFieldConfigProperty.ShowThumbnails`.
	ShowThumbnails interface{} `field:"optional" json:"showThumbnails" yaml:"showThumbnails"`
}

