package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2FlowLog AWS CloudFormation Resource (AWS::EC2::FlowLog)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html
type AWSEC2FlowLog struct {

	// DeliverLogsPermissionArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-deliverlogspermissionarn
	DeliverLogsPermissionArn *stringIntrinsic `json:"DeliverLogsPermissionArn,omitempty"`

	// LogGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-loggroupname
	LogGroupName *stringIntrinsic `json:"LogGroupName,omitempty"`

	// ResourceId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourceid
	ResourceId *stringIntrinsic `json:"ResourceId,omitempty"`

	// ResourceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-resourcetype
	ResourceType *stringIntrinsic `json:"ResourceType,omitempty"`

	// TrafficType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-flowlog.html#cfn-ec2-flowlog-traffictype
	TrafficType *stringIntrinsic `json:"TrafficType,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2FlowLog) AWSCloudFormationType() string {
	return "AWS::EC2::FlowLog"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2FlowLog) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2FlowLog
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEC2FlowLog) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2FlowLog
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSEC2FlowLog(*res.Properties)
	}

	return nil
}

// GetAllAWSEC2FlowLogResources retrieves all AWSEC2FlowLog items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2FlowLogResources() map[string]AWSEC2FlowLog {
	results := map[string]AWSEC2FlowLog{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2FlowLog:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::FlowLog" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2FlowLog
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEC2FlowLogWithName retrieves all AWSEC2FlowLog items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2FlowLogWithName(name string) (AWSEC2FlowLog, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2FlowLog:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::FlowLog" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSEC2FlowLog
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSEC2FlowLog{}, errors.New("resource not found")
}
