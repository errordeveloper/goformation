package cloudformation

// AWSDMSEndpoint_S3Settings AWS CloudFormation Resource (AWS::DMS::Endpoint.S3Settings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html
type AWSDMSEndpoint_S3Settings struct {

	// BucketFolder AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-bucketfolder
	BucketFolder *stringIntrinsic `json:"BucketFolder,omitempty"`

	// BucketName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-bucketname
	BucketName *stringIntrinsic `json:"BucketName,omitempty"`

	// CompressionType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-compressiontype
	CompressionType *stringIntrinsic `json:"CompressionType,omitempty"`

	// CsvDelimiter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-csvdelimiter
	CsvDelimiter *stringIntrinsic `json:"CsvDelimiter,omitempty"`

	// CsvRowDelimiter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-csvrowdelimiter
	CsvRowDelimiter *stringIntrinsic `json:"CsvRowDelimiter,omitempty"`

	// ExternalTableDefinition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-externaltabledefinition
	ExternalTableDefinition *stringIntrinsic `json:"ExternalTableDefinition,omitempty"`

	// ServiceAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-serviceaccessrolearn
	ServiceAccessRoleArn *stringIntrinsic `json:"ServiceAccessRoleArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSEndpoint_S3Settings) AWSCloudFormationType() string {
	return "AWS::DMS::Endpoint.S3Settings"
}
