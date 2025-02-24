// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = ackv1alpha1.AWSAccountID("")
)

// An object that contains details about when a principal in the reported Organizations
// entity last attempted to access an Amazon Web Services service. A principal
// can be an IAM user, an IAM role, or the Amazon Web Services account root
// user within the reported Organizations entity.
//
// This data type is a response element in the GetOrganizationsAccessReport
// operation.
type AccessDetail struct {
	LastAuthenticatedTime *metav1.Time `json:"lastAuthenticatedTime,omitempty"`
	Region                *string      `json:"region,omitempty"`
}

// Contains information about an Amazon Web Services access key.
//
// This data type is used as a response element in the CreateAccessKey and ListAccessKeys
// operations.
//
// The SecretAccessKey value is returned only in response to CreateAccessKey.
// You can get a secret access key only when you first create an access key;
// you cannot recover the secret access key later. If you lose a secret access
// key, you must create a new access key.
type AccessKey struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`
	UserName   *string      `json:"userName,omitempty"`
}

// Contains information about the last time an Amazon Web Services access key
// was used since IAM began tracking this information on April 22, 2015.
//
// This data type is used as a response element in the GetAccessKeyLastUsed
// operation.
type AccessKeyLastUsed struct {
	LastUsedDate *metav1.Time `json:"lastUsedDate,omitempty"`
	Region       *string      `json:"region,omitempty"`
	ServiceName  *string      `json:"serviceName,omitempty"`
}

// Contains information about an Amazon Web Services access key, without its
// secret key.
//
// This data type is used as a response element in the ListAccessKeys operation.
type AccessKeyMetadata struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`
	UserName   *string      `json:"userName,omitempty"`
}

// Contains information about an attached permissions boundary.
//
// An attached permissions boundary is a managed policy that has been attached
// to a user or role to set the permissions boundary.
//
// For more information about permissions boundaries, see Permissions boundaries
// for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
// in the IAM User Guide.
type AttachedPermissionsBoundary struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	PermissionsBoundaryARN  *string `json:"permissionsBoundaryARN,omitempty"`
	PermissionsBoundaryType *string `json:"permissionsBoundaryType,omitempty"`
}

// Contains information about an attached policy.
//
// An attached policy is a managed policy that has been attached to a user,
// group, or role. This data type is used as a response element in the ListAttachedGroupPolicies,
// ListAttachedRolePolicies, ListAttachedUserPolicies, and GetAccountAuthorizationDetails
// operations.
//
// For more information about managed policies, refer to Managed policies and
// inline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
type AttachedPolicy struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	PolicyARN  *string `json:"policyARN,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
}

// An object that contains details about when the IAM entities (users or roles)
// were last used in an attempt to access the specified Amazon Web Services
// service.
//
// This data type is a response element in the GetServiceLastAccessedDetailsWithEntities
// operation.
type EntityDetails struct {
	LastAuthenticated *metav1.Time `json:"lastAuthenticated,omitempty"`
}

// Contains details about the specified entity (user or role).
//
// This data type is an element of the EntityDetails object.
type EntityInfo struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN  *string `json:"arn,omitempty"`
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
}

// Contains information about the reason that the operation failed.
//
// This data type is used as a response element in the GetOrganizationsAccessReport,
// GetServiceLastAccessedDetails, and GetServiceLastAccessedDetailsWithEntities
// operations.
type ErrorDetails struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Contains information about an IAM group, including all of the group's policies.
//
// This data type is used as a response element in the GetAccountAuthorizationDetails
// operation.
type GroupDetail struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN        *string      `json:"arn,omitempty"`
	CreateDate *metav1.Time `json:"createDate,omitempty"`
	GroupID    *string      `json:"groupID,omitempty"`
	GroupName  *string      `json:"groupName,omitempty"`
	Path       *string      `json:"path,omitempty"`
}

// Contains information about an IAM group entity.
//
// This data type is used as a response element in the following operations:
//
//    * CreateGroup
//
//    * GetGroup
//
//    * ListGroups
type Group_SDK struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN        *string      `json:"arn,omitempty"`
	CreateDate *metav1.Time `json:"createDate,omitempty"`
	GroupID    *string      `json:"groupID,omitempty"`
	GroupName  *string      `json:"groupName,omitempty"`
	Path       *string      `json:"path,omitempty"`
}

// Contains information about an instance profile.
//
// This data type is used as a response element in the following operations:
//
//    * CreateInstanceProfile
//
//    * GetInstanceProfile
//
//    * ListInstanceProfiles
//
//    * ListInstanceProfilesForRole
type InstanceProfile struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN               *string      `json:"arn,omitempty"`
	CreateDate        *metav1.Time `json:"createDate,omitempty"`
	InstanceProfileID *string      `json:"instanceProfileID,omitempty"`
	Path              *string      `json:"path,omitempty"`
	// Contains a list of IAM roles.
	//
	// This data type is used as a response element in the ListRoles operation.
	Roles []*Role_SDK `json:"roles,omitempty"`
	Tags  []*Tag      `json:"tags,omitempty"`
}

// Contains the user name and password create date for a user.
//
// This data type is used as a response element in the CreateLoginProfile and
// GetLoginProfile operations.
type LoginProfile struct {
	CreateDate            *metav1.Time `json:"createDate,omitempty"`
	PasswordResetRequired *bool        `json:"passwordResetRequired,omitempty"`
	UserName              *string      `json:"userName,omitempty"`
}

// Contains information about an MFA device.
//
// This data type is used as a response element in the ListMFADevices operation.
type MFADevice struct {
	EnableDate *metav1.Time `json:"enableDate,omitempty"`
	UserName   *string      `json:"userName,omitempty"`
}

// Contains information about a managed policy, including the policy's ARN,
// versions, and the number of principal entities (users, groups, and roles)
// that the policy is attached to.
//
// This data type is used as a response element in the GetAccountAuthorizationDetails
// operation.
//
// For more information about managed policies, see Managed policies and inline
// policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
type ManagedPolicyDetail struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN                           *string      `json:"arn,omitempty"`
	AttachmentCount               *int64       `json:"attachmentCount,omitempty"`
	CreateDate                    *metav1.Time `json:"createDate,omitempty"`
	DefaultVersionID              *string      `json:"defaultVersionID,omitempty"`
	Description                   *string      `json:"description,omitempty"`
	IsAttachable                  *bool        `json:"isAttachable,omitempty"`
	Path                          *string      `json:"path,omitempty"`
	PermissionsBoundaryUsageCount *int64       `json:"permissionsBoundaryUsageCount,omitempty"`
	PolicyID                      *string      `json:"policyID,omitempty"`
	PolicyName                    *string      `json:"policyName,omitempty"`
	UpdateDate                    *metav1.Time `json:"updateDate,omitempty"`
}

// Contains the Amazon Resource Name (ARN) for an IAM OpenID Connect provider.
type OpenIDConnectProviderListEntry struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN *string `json:"arn,omitempty"`
}

// Contains information about the effect that Organizations has on a policy
// simulation.
type OrganizationsDecisionDetail struct {
	AllowedByOrganizations *bool `json:"allowedByOrganizations,omitempty"`
}

// Contains information about the account password policy.
//
// This data type is used as a response element in the GetAccountPasswordPolicy
// operation.
type PasswordPolicy struct {
	AllowUsersToChangePassword *bool `json:"allowUsersToChangePassword,omitempty"`
	ExpirePasswords            *bool `json:"expirePasswords,omitempty"`
	RequireLowercaseCharacters *bool `json:"requireLowercaseCharacters,omitempty"`
	RequireNumbers             *bool `json:"requireNumbers,omitempty"`
	RequireSymbols             *bool `json:"requireSymbols,omitempty"`
	RequireUppercaseCharacters *bool `json:"requireUppercaseCharacters,omitempty"`
}

// Contains information about the effect that a permissions boundary has on
// a policy simulation when the boundary is applied to an IAM entity.
type PermissionsBoundaryDecisionDetail struct {
	AllowedByPermissionsBoundary *bool `json:"allowedByPermissionsBoundary,omitempty"`
}

// Contains information about an IAM policy, including the policy document.
//
// This data type is used as a response element in the GetAccountAuthorizationDetails
// operation.
type PolicyDetail struct {
	PolicyDocument *string `json:"policyDocument,omitempty"`
	PolicyName     *string `json:"policyName,omitempty"`
}

// Contains details about the permissions policies that are attached to the
// specified identity (user, group, or role).
//
// This data type is an element of the ListPoliciesGrantingServiceAccessEntry
// object.
type PolicyGrantingServiceAccess struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	PolicyARN  *string `json:"policyARN,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
}

// Contains information about a group that a managed policy is attached to.
//
// This data type is used as a response element in the ListEntitiesForPolicy
// operation.
//
// For more information about managed policies, refer to Managed policies and
// inline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
type PolicyGroup struct {
	GroupID   *string `json:"groupID,omitempty"`
	GroupName *string `json:"groupName,omitempty"`
}

// Contains information about a role that a managed policy is attached to.
//
// This data type is used as a response element in the ListEntitiesForPolicy
// operation.
//
// For more information about managed policies, refer to Managed policies and
// inline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
type PolicyRole struct {
	RoleID   *string `json:"roleID,omitempty"`
	RoleName *string `json:"roleName,omitempty"`
}

// Contains information about a user that a managed policy is attached to.
//
// This data type is used as a response element in the ListEntitiesForPolicy
// operation.
//
// For more information about managed policies, refer to Managed policies and
// inline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
type PolicyUser struct {
	UserID   *string `json:"userID,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// Contains information about a version of a managed policy.
//
// This data type is used as a response element in the CreatePolicyVersion,
// GetPolicyVersion, ListPolicyVersions, and GetAccountAuthorizationDetails
// operations.
//
// For more information about managed policies, refer to Managed policies and
// inline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
type PolicyVersion struct {
	CreateDate       *metav1.Time `json:"createDate,omitempty"`
	Document         *string      `json:"document,omitempty"`
	IsDefaultVersion *bool        `json:"isDefaultVersion,omitempty"`
	VersionID        *string      `json:"versionID,omitempty"`
}

// Contains information about a managed policy.
//
// This data type is used as a response element in the CreatePolicy, GetPolicy,
// and ListPolicies operations.
//
// For more information about managed policies, refer to Managed policies and
// inline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
type Policy_SDK struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN                           *string      `json:"arn,omitempty"`
	AttachmentCount               *int64       `json:"attachmentCount,omitempty"`
	CreateDate                    *metav1.Time `json:"createDate,omitempty"`
	DefaultVersionID              *string      `json:"defaultVersionID,omitempty"`
	Description                   *string      `json:"description,omitempty"`
	IsAttachable                  *bool        `json:"isAttachable,omitempty"`
	Path                          *string      `json:"path,omitempty"`
	PermissionsBoundaryUsageCount *int64       `json:"permissionsBoundaryUsageCount,omitempty"`
	PolicyID                      *string      `json:"policyID,omitempty"`
	PolicyName                    *string      `json:"policyName,omitempty"`
	Tags                          []*Tag       `json:"tags,omitempty"`
	UpdateDate                    *metav1.Time `json:"updateDate,omitempty"`
}

// Contains information about an IAM role, including all of the role's policies.
//
// This data type is used as a response element in the GetAccountAuthorizationDetails
// operation.
type RoleDetail struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN                      *string      `json:"arn,omitempty"`
	AssumeRolePolicyDocument *string      `json:"assumeRolePolicyDocument,omitempty"`
	CreateDate               *metav1.Time `json:"createDate,omitempty"`
	Path                     *string      `json:"path,omitempty"`
	// Contains information about an attached permissions boundary.
	//
	// An attached permissions boundary is a managed policy that has been attached
	// to a user or role to set the permissions boundary.
	//
	// For more information about permissions boundaries, see Permissions boundaries
	// for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide.
	PermissionsBoundary *AttachedPermissionsBoundary `json:"permissionsBoundary,omitempty"`
	RoleID              *string                      `json:"roleID,omitempty"`
	// Contains information about the last time that an IAM role was used. This
	// includes the date and time and the Region in which the role was last used.
	// Activity is only reported for the trailing 400 days. This period can be shorter
	// if your Region began supporting these features within the last year. The
	// role might have been used more than 400 days ago. For more information, see
	// Regions where data is tracked (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period)
	// in the IAM User Guide.
	//
	// This data type is returned as a response element in the GetRole and GetAccountAuthorizationDetails
	// operations.
	RoleLastUsed *RoleLastUsed `json:"roleLastUsed,omitempty"`
	RoleName     *string       `json:"roleName,omitempty"`
	Tags         []*Tag        `json:"tags,omitempty"`
}

// Contains information about the last time that an IAM role was used. This
// includes the date and time and the Region in which the role was last used.
// Activity is only reported for the trailing 400 days. This period can be shorter
// if your Region began supporting these features within the last year. The
// role might have been used more than 400 days ago. For more information, see
// Regions where data is tracked (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period)
// in the IAM User Guide.
//
// This data type is returned as a response element in the GetRole and GetAccountAuthorizationDetails
// operations.
type RoleLastUsed struct {
	LastUsedDate *metav1.Time `json:"lastUsedDate,omitempty"`
	Region       *string      `json:"region,omitempty"`
}

// Contains information about an IAM role. This structure is returned as a response
// element in several API operations that interact with roles.
type Role_SDK struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN                      *string      `json:"arn,omitempty"`
	AssumeRolePolicyDocument *string      `json:"assumeRolePolicyDocument,omitempty"`
	CreateDate               *metav1.Time `json:"createDate,omitempty"`
	Description              *string      `json:"description,omitempty"`
	MaxSessionDuration       *int64       `json:"maxSessionDuration,omitempty"`
	Path                     *string      `json:"path,omitempty"`
	// Contains information about an attached permissions boundary.
	//
	// An attached permissions boundary is a managed policy that has been attached
	// to a user or role to set the permissions boundary.
	//
	// For more information about permissions boundaries, see Permissions boundaries
	// for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide.
	PermissionsBoundary *AttachedPermissionsBoundary `json:"permissionsBoundary,omitempty"`
	RoleID              *string                      `json:"roleID,omitempty"`
	// Contains information about the last time that an IAM role was used. This
	// includes the date and time and the Region in which the role was last used.
	// Activity is only reported for the trailing 400 days. This period can be shorter
	// if your Region began supporting these features within the last year. The
	// role might have been used more than 400 days ago. For more information, see
	// Regions where data is tracked (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period)
	// in the IAM User Guide.
	//
	// This data type is returned as a response element in the GetRole and GetAccountAuthorizationDetails
	// operations.
	RoleLastUsed *RoleLastUsed `json:"roleLastUsed,omitempty"`
	RoleName     *string       `json:"roleName,omitempty"`
	Tags         []*Tag        `json:"tags,omitempty"`
}

// Contains the list of SAML providers for this account.
type SAMLProviderListEntry struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN        *string      `json:"arn,omitempty"`
	CreateDate *metav1.Time `json:"createDate,omitempty"`
	ValidUntil *metav1.Time `json:"validUntil,omitempty"`
}

// Contains information about an SSH public key.
//
// This data type is used as a response element in the GetSSHPublicKey and UploadSSHPublicKey
// operations.
type SSHPublicKey struct {
	UploadDate *metav1.Time `json:"uploadDate,omitempty"`
	UserName   *string      `json:"userName,omitempty"`
}

// Contains information about an SSH public key, without the key's body or fingerprint.
//
// This data type is used as a response element in the ListSSHPublicKeys operation.
type SSHPublicKeyMetadata struct {
	UploadDate *metav1.Time `json:"uploadDate,omitempty"`
	UserName   *string      `json:"userName,omitempty"`
}

// Contains information about a server certificate.
//
// This data type is used as a response element in the GetServerCertificate
// operation.
type ServerCertificate struct {
	Tags []*Tag `json:"tags,omitempty"`
}

// Contains information about a server certificate without its certificate body,
// certificate chain, and private key.
//
// This data type is used as a response element in the UploadServerCertificate
// and ListServerCertificates operations.
type ServerCertificateMetadata struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN                 *string      `json:"arn,omitempty"`
	Expiration          *metav1.Time `json:"expiration,omitempty"`
	Path                *string      `json:"path,omitempty"`
	ServerCertificateID *string      `json:"serverCertificateID,omitempty"`
	UploadDate          *metav1.Time `json:"uploadDate,omitempty"`
}

// Contains details about the most recent attempt to access the service.
//
// This data type is used as a response element in the GetServiceLastAccessedDetails
// operation.
type ServiceLastAccessed struct {
	LastAuthenticated *metav1.Time `json:"lastAuthenticated,omitempty"`
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	LastAuthenticatedEntity *string `json:"lastAuthenticatedEntity,omitempty"`
	LastAuthenticatedRegion *string `json:"lastAuthenticatedRegion,omitempty"`
}

// Contains the details of a service-specific credential.
type ServiceSpecificCredential struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`
	UserName   *string      `json:"userName,omitempty"`
}

// Contains additional details about a service-specific credential.
type ServiceSpecificCredentialMetadata struct {
	CreateDate *metav1.Time `json:"createDate,omitempty"`
	UserName   *string      `json:"userName,omitempty"`
}

// Contains information about an X.509 signing certificate.
//
// This data type is used as a response element in the UploadSigningCertificate
// and ListSigningCertificates operations.
type SigningCertificate struct {
	UploadDate *metav1.Time `json:"uploadDate,omitempty"`
	UserName   *string      `json:"userName,omitempty"`
}

// A structure that represents user-provided metadata that can be associated
// with an IAM resource. For more information about tagging, see Tagging IAM
// resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
// in the IAM User Guide.
type Tag struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Contains details about the most recent attempt to access an action within
// the service.
//
// This data type is used as a response element in the GetServiceLastAccessedDetails
// operation.
type TrackedActionLastAccessed struct {
	ActionName *string `json:"actionName,omitempty"`
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	LastAccessedEntity *string      `json:"lastAccessedEntity,omitempty"`
	LastAccessedRegion *string      `json:"lastAccessedRegion,omitempty"`
	LastAccessedTime   *metav1.Time `json:"lastAccessedTime,omitempty"`
}

// Contains information about an IAM user entity.
//
// This data type is used as a response element in the following operations:
//
//    * CreateUser
//
//    * GetUser
//
//    * ListUsers
type User struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN              *string      `json:"arn,omitempty"`
	CreateDate       *metav1.Time `json:"createDate,omitempty"`
	PasswordLastUsed *metav1.Time `json:"passwordLastUsed,omitempty"`
	Path             *string      `json:"path,omitempty"`
	// Contains information about an attached permissions boundary.
	//
	// An attached permissions boundary is a managed policy that has been attached
	// to a user or role to set the permissions boundary.
	//
	// For more information about permissions boundaries, see Permissions boundaries
	// for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide.
	PermissionsBoundary *AttachedPermissionsBoundary `json:"permissionsBoundary,omitempty"`
	Tags                []*Tag                       `json:"tags,omitempty"`
	UserID              *string                      `json:"userID,omitempty"`
	UserName            *string                      `json:"userName,omitempty"`
}

// Contains information about an IAM user, including all the user's policies
// and all the IAM groups the user is in.
//
// This data type is used as a response element in the GetAccountAuthorizationDetails
// operation.
type UserDetail struct {
	// The Amazon Resource Name (ARN). ARNs are unique identifiers for Amazon Web
	// Services resources.
	//
	// For more information about ARNs, go to Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	ARN        *string      `json:"arn,omitempty"`
	CreateDate *metav1.Time `json:"createDate,omitempty"`
	Path       *string      `json:"path,omitempty"`
	// Contains information about an attached permissions boundary.
	//
	// An attached permissions boundary is a managed policy that has been attached
	// to a user or role to set the permissions boundary.
	//
	// For more information about permissions boundaries, see Permissions boundaries
	// for IAM identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide.
	PermissionsBoundary *AttachedPermissionsBoundary `json:"permissionsBoundary,omitempty"`
	Tags                []*Tag                       `json:"tags,omitempty"`
	UserID              *string                      `json:"userID,omitempty"`
	UserName            *string                      `json:"userName,omitempty"`
}

// Contains information about a virtual MFA device.
type VirtualMFADevice struct {
	EnableDate *metav1.Time `json:"enableDate,omitempty"`
	Tags       []*Tag       `json:"tags,omitempty"`
	// Contains information about an IAM user entity.
	//
	// This data type is used as a response element in the following operations:
	//
	//    * CreateUser
	//
	//    * GetUser
	//
	//    * ListUsers
	User *User `json:"user,omitempty"`
}
