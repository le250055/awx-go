package awx

// TODO Add mock data and then enable test

import (
	"testing"
	"time"
)

func TestListCredentials(t *testing.T) {
	var (
		expectListCredentialsResponse = []*Credential{
			{
				ID:   1,
				Type: "credential",
				URL:  "/api/v2/credentials/1/",
				Related: &Related{
					ObjectRoles:    "/api/v2/credentials/1/object_roles/",
					AccessList:     "/api/v2/credentials/1/access_list/",
					CredentialType: "/api/v2/credential_types/1/",
					CreatedBy:      "/api/v2/users/1/",
					ModifiedBy:     "/api/v2/users/1/",
					OwnerUsers:     "/api/v2/credentials/1/owner_users/",
					OwnerTeams:     "/api/v2/credentials/1/owner_teams/",
					Copy:           "/api/v2/credentials/1/copy/",
					ActivityStream: "/api/v2/credentials/1/activity_stream/",
				},
				SummaryFields: &Summary{
					CredentialType: &CredentialTypeSummary{
						ID:          1,
						Name:        "Machine",
						Description: "",
					},
					CreatedBy: &ByUserSummary{
						ID:        1,
						Username:  "admin",
						FirstName: "",
						LastName:  "",
					},
					ModifiedBy: &ByUserSummary{
						ID:        1,
						Username:  "admin",
						FirstName: "",
						LastName:  "",
					},
					ObjectRoles: &ObjectRoles{
						AdminRole: &ObjectRole{
							ID:          19,
							Description: "Can manage all aspects of the credential",
							Name:        "Admin",
						},
						UseRole: &ObjectRole{
							ID:          20,
							Description: "Can use the credential in a job template",
							Name:        "Use",
						},
						ReadRole: &ObjectRole{
							ID:          21,
							Description: "May view settings for the credential",
							Name:        "Read",
						},
					},
					UserCapabilities: &UserCapabilities{
						Edit:   true,
						Delete: true,
						Copy:   true,
					},
					Owners: []OwnerUserSummary{
						{
							ID:          1,
							Type:        "user",
							Name:        "admin",
							URL:         "/api/v2/users/1/",
							Description: "",
						},
					},
				},
				Created: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2021-01-11T17:37:27.381945Z")
					return t
				}(),
				Modified: func() time.Time {
					t, _ := time.Parse(time.RFC3339, "2021-01-11T17:37:27.381975Z")
					return t
				}(),
				Name:           "Demo Credential",
				Description:    "",
				CredentialType: 1,
				Inputs: map[string]interface{}{
					"username": "admin",
				},
				Kind:       "ssh",
				Cloud:      false,
				Kubernetes: false,
			},
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, _, err := awx.CredentialService.ListCredentials(map[string]string{
		"name": "Demo Credential",
	})

	if err != nil {
		t.Fatalf("ListCredentials err: %s", err)
	} else {
		checkAPICallResult(t, expectListCredentialsResponse, result)
		t.Log("ListCredentials passed!")
	}
}

func TestCreateCredentials(t *testing.T) {
	var (
		expectCreateCredentialsResponse = &Credential{
			ID:   1,
			Type: "credential",
			URL:  "/api/v2/credentials/1/",
			Related: &Related{
				ObjectRoles:    "/api/v2/credentials/1/object_roles/",
				AccessList:     "/api/v2/credentials/1/access_list/",
				CredentialType: "/api/v2/credential_types/1/",
				CreatedBy:      "/api/v2/users/1/",
				ModifiedBy:     "/api/v2/users/1/",
				OwnerUsers:     "/api/v2/credentials/1/owner_users/",
				OwnerTeams:     "/api/v2/credentials/1/owner_teams/",
				Copy:           "/api/v2/credentials/1/copy/",
				ActivityStream: "/api/v2/credentials/1/activity_stream/",
			},
			SummaryFields: &Summary{
				CredentialType: &CredentialTypeSummary{
					ID:          1,
					Name:        "Machine",
					Description: "",
				},
				CreatedBy: &ByUserSummary{
					ID:        1,
					Username:  "admin",
					FirstName: "",
					LastName:  "",
				},
				ModifiedBy: &ByUserSummary{
					ID:        1,
					Username:  "admin",
					FirstName: "",
					LastName:  "",
				},
				ObjectRoles: &ObjectRoles{
					AdminRole: &ObjectRole{
						ID:          19,
						Description: "Can manage all aspects of the credential",
						Name:        "Admin",
					},
					UseRole: &ObjectRole{
						ID:          20,
						Description: "Can use the credential in a job template",
						Name:        "Use",
					},
					ReadRole: &ObjectRole{
						ID:          21,
						Description: "May view settings for the credential",
						Name:        "Read",
					},
				},
				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Delete: true,
					Copy:   true,
				},
				Owners: []OwnerUserSummary{
					{
						ID:          1,
						Type:        "user",
						Name:        "admin",
						URL:         "/api/v2/users/1/",
						Description: "",
					},
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2021-01-11T17:37:27.381945Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2021-01-11T17:37:27.381975Z")
				return t
			}(),
			Name:           "Demo Credential",
			Description:    "",
			CredentialType: 1,
			Inputs: map[string]interface{}{
				"username": "admin",
			},
			Kind:       "ssh",
			Cloud:      false,
			Kubernetes: false,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.CredentialService.CreateCredential(map[string]interface{}{
		"name":            "Demo Credential",
		"organization":    1,
		"credential_type": 1,
	}, map[string]string{})
	if err != nil {
		t.Fatalf("CreateCredentials err: %s", err)
	} else {
		checkAPICallResult(t, expectCreateCredentialsResponse, result)
		t.Log("CreateCredentials passed!")
	}
}

func TestUpdateCredentials(t *testing.T) {
	var (
		expectUpdateCredentialsResponse = &Credential{
			ID:   1,
			Type: "credential",
			URL:  "/api/v2/credentials/1/",
			Related: &Related{
				ObjectRoles:    "/api/v2/credentials/1/object_roles/",
				AccessList:     "/api/v2/credentials/1/access_list/",
				CredentialType: "/api/v2/credential_types/1/",
				CreatedBy:      "/api/v2/users/1/",
				ModifiedBy:     "/api/v2/users/1/",
				OwnerUsers:     "/api/v2/credentials/1/owner_users/",
				OwnerTeams:     "/api/v2/credentials/1/owner_teams/",
				Copy:           "/api/v2/credentials/1/copy/",
				ActivityStream: "/api/v2/credentials/1/activity_stream/",
			},
			SummaryFields: &Summary{
				CredentialType: &CredentialTypeSummary{
					ID:          1,
					Name:        "Machine",
					Description: "",
				},
				CreatedBy: &ByUserSummary{
					ID:        1,
					Username:  "admin",
					FirstName: "",
					LastName:  "",
				},
				ModifiedBy: &ByUserSummary{
					ID:        1,
					Username:  "admin",
					FirstName: "",
					LastName:  "",
				},
				ObjectRoles: &ObjectRoles{
					AdminRole: &ObjectRole{
						ID:          19,
						Description: "Can manage all aspects of the credential",
						Name:        "Admin",
					},
					UseRole: &ObjectRole{
						ID:          20,
						Description: "Can use the credential in a job template",
						Name:        "Use",
					},
					ReadRole: &ObjectRole{
						ID:          21,
						Description: "May view settings for the credential",
						Name:        "Read",
					},
				},
				UserCapabilities: &UserCapabilities{
					Edit:   true,
					Delete: true,
					Copy:   true,
				},
				Owners: []OwnerUserSummary{
					{
						ID:          1,
						Type:        "user",
						Name:        "admin",
						URL:         "/api/v2/users/1/",
						Description: "",
					},
				},
			},
			Created: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2021-01-11T17:37:27.381945Z")
				return t
			}(),
			Modified: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2021-01-11T17:37:27.381975Z")
				return t
			}(),
			Name:           "Demo Credential",
			Description:    "Demo Credential Updated",
			CredentialType: 1,
			Inputs: map[string]interface{}{
				"username": "admin",
			},
			Kind:       "ssh",
			Cloud:      false,
			Kubernetes: false,
		}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.CredentialService.UpdateCredential(1, map[string]interface{}{
		"name":        "Demo Credential",
		"description": "Demo Credential Updated",
	}, map[string]string{})
	if err != nil {
		t.Fatalf("TestUpdateCredentials err: %s", err)
	} else {
		checkAPICallResult(t, expectUpdateCredentialsResponse, result)
		t.Log("TestUpdateCredentials passed!")
	}
}

func TestDeleteCredentials(t *testing.T) {
	var (
		expectDeleteCredentialsResponse = &Credential{}
	)

	awx := NewAWX(testAwxHost, testAwxUserName, testAwxPasswd, nil)
	result, err := awx.CredentialService.DeleteCredential(1)

	if err != nil {
		t.Fatalf("DeleteCredentials err: %s", err)
	} else {
		checkAPICallResult(t, expectDeleteCredentialsResponse, result)
		t.Log("DeleteCredentials passed!")
	}
}
