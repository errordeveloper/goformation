package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSIAMInstanceProfile AWS CloudFormation Resource (AWS::IAM::InstanceProfile)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
type AWSIAMInstanceProfile struct {

	// InstanceProfileName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-instanceprofilename
	InstanceProfileName *Value `json:"InstanceProfileName,omitempty"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-path
	Path *Value `json:"Path,omitempty"`

	// Roles AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-roles
	Roles []*Value `json:"Roles,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMInstanceProfile) AWSCloudFormationType() string {
	return "AWS::IAM::InstanceProfile"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSIAMInstanceProfile) MarshalJSON() ([]byte, error) {
	type Properties AWSIAMInstanceProfile
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
func (r *AWSIAMInstanceProfile) UnmarshalJSON(b []byte) error {
	type Properties AWSIAMInstanceProfile
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
		*r = AWSIAMInstanceProfile(*res.Properties)
	}

	return nil
}

// GetAllAWSIAMInstanceProfileResources retrieves all AWSIAMInstanceProfile items from an AWS CloudFormation template
func (t *Template) GetAllAWSIAMInstanceProfileResources() map[string]AWSIAMInstanceProfile {
	results := map[string]AWSIAMInstanceProfile{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSIAMInstanceProfile:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::IAM::InstanceProfile" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSIAMInstanceProfile{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSIAMInstanceProfileWithName retrieves all AWSIAMInstanceProfile items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMInstanceProfileWithName(name string) (AWSIAMInstanceProfile, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSIAMInstanceProfile:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::IAM::InstanceProfile" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSIAMInstanceProfile{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSIAMInstanceProfile{}, errors.New("resource not found")
}
