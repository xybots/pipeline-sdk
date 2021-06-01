/*
 * Pipeline API
 *
 * Pipeline is a feature rich application platform, built for containers on top of Kubernetes to automate the DevOps experience, continuous application development and the lifecycle of deployments. 
 *
 * API version: latest
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline
// AmazonAutoScalingGroup struct for AmazonAutoScalingGroup
type AmazonAutoScalingGroup struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Zones []string `json:"zones"`
	InstanceType string `json:"instanceType"`
	LaunchConfigurationName string `json:"launchConfigurationName"`
	LaunchTemplate map[string]interface{} `json:"launchTemplate,omitempty"`
	VpcID string `json:"vpcID"`
	SecurityGroupID string `json:"securityGroupID"`
	Subnets []string `json:"subnets"`
	Tags map[string]map[string]interface{} `json:"tags"`
	SpotPrice string `json:"spotPrice"`
	Size AmazonAutoScalingGroupSize `json:"size"`
}
