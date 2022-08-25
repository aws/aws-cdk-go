// Version 2 of the AWS Cloud Development Kit library
package awscdk


type TagType string

const (
	TagType_STANDARD TagType = "STANDARD"
	TagType_AUTOSCALING_GROUP TagType = "AUTOSCALING_GROUP"
	TagType_MAP TagType = "MAP"
	TagType_KEY_VALUE TagType = "KEY_VALUE"
	TagType_NOT_TAGGABLE TagType = "NOT_TAGGABLE"
)

