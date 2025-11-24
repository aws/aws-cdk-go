package mixinsawsamplifyuibuilder


// Describes the configuration for the file uploader field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fileUploaderFieldConfigProperty := &FileUploaderFieldConfigProperty{
//   	AcceptedFileTypes: []*string{
//   		jsii.String("acceptedFileTypes"),
//   	},
//   	AccessLevel: jsii.String("accessLevel"),
//   	IsResumable: jsii.Boolean(false),
//   	MaxFileCount: jsii.Number(123),
//   	MaxSize: jsii.Number(123),
//   	ShowThumbnails: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fileuploaderfieldconfig.html
//
type CfnFormPropsMixin_FileUploaderFieldConfigProperty struct {
	// The file types that are allowed to be uploaded by the file uploader.
	//
	// Provide this information in an array of strings specifying the valid file extensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fileuploaderfieldconfig.html#cfn-amplifyuibuilder-form-fileuploaderfieldconfig-acceptedfiletypes
	//
	AcceptedFileTypes *[]*string `field:"optional" json:"acceptedFileTypes" yaml:"acceptedFileTypes"`
	// The access level to assign to the uploaded files in the Amazon S3 bucket where they are stored.
	//
	// The valid values for this property are `private` , `protected` , or `public` . For detailed information about the permissions associated with each access level, see [File access levels](https://docs.aws.amazon.com/https://docs.amplify.aws/lib/storage/configureaccess/q/platform/js/) in the *Amplify documentation* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fileuploaderfieldconfig.html#cfn-amplifyuibuilder-form-fileuploaderfieldconfig-accesslevel
	//
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// Allows the file upload operation to be paused and resumed. The default value is `false` .
	//
	// When `isResumable` is set to `true` , the file uploader uses a multipart upload to break the files into chunks before upload. The progress of the upload isn't continuous, because the file uploader uploads a chunk at a time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fileuploaderfieldconfig.html#cfn-amplifyuibuilder-form-fileuploaderfieldconfig-isresumable
	//
	IsResumable interface{} `field:"optional" json:"isResumable" yaml:"isResumable"`
	// Specifies the maximum number of files that can be selected to upload.
	//
	// The default value is an unlimited number of files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fileuploaderfieldconfig.html#cfn-amplifyuibuilder-form-fileuploaderfieldconfig-maxfilecount
	//
	MaxFileCount *float64 `field:"optional" json:"maxFileCount" yaml:"maxFileCount"`
	// The maximum file size in bytes that the file uploader will accept.
	//
	// The default value is an unlimited file size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fileuploaderfieldconfig.html#cfn-amplifyuibuilder-form-fileuploaderfieldconfig-maxsize
	//
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Specifies whether to display or hide the image preview after selecting a file for upload.
	//
	// The default value is `true` to display the image preview.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amplifyuibuilder-form-fileuploaderfieldconfig.html#cfn-amplifyuibuilder-form-fileuploaderfieldconfig-showthumbnails
	//
	ShowThumbnails interface{} `field:"optional" json:"showThumbnails" yaml:"showThumbnails"`
}

