package mockdata

var (
	// MockedListCredentialsResponse mocked `ListCredentials` endpoint response
	MockedListCredentialsResponse = []byte(`
		{
			"count": 1,
			"next": null,
			"previous": null,
			"results": [
				{
					"id": 1,
					"type": "credential",
					"url": "/api/v2/credentials/1/",
					"related": {
						"created_by": "/api/v2/users/1/",
						"modified_by": "/api/v2/users/1/",
						"activity_stream": "/api/v2/credentials/1/activity_stream/",
						"access_list": "/api/v2/credentials/1/access_list/",
						"object_roles": "/api/v2/credentials/1/object_roles/",
						"owner_users": "/api/v2/credentials/1/owner_users/",
						"owner_teams": "/api/v2/credentials/1/owner_teams/",
						"copy": "/api/v2/credentials/1/copy/",
						"input_sources": "/api/v2/credentials/1/input_sources/",
						"credential_type": "/api/v2/credential_types/1/",
						"user": "/api/v2/users/1/"
					},
					"summary_fields": {
						"credential_type": {
							"id": 1,
							"name": "Machine",
							"description": ""
						},
						"created_by": {
							"id": 1,
							"username": "admin",
							"first_name": "",
							"last_name": ""
						},
						"modified_by": {
							"id": 1,
							"username": "admin",
							"first_name": "",
							"last_name": ""
						},
						"object_roles": {
							"admin_role": {
								"description": "Can manage all aspects of the credential",
								"name": "Admin",
								"id": 19
							},
							"use_role": {
								"description": "Can use the credential in a job template",
								"name": "Use",
								"id": 20
							},
							"read_role": {
								"description": "May view settings for the credential",
								"name": "Read",
								"id": 21
							}
						},
						"user_capabilities": {
							"edit": true,
							"delete": true,
							"copy": true,
							"use": true
						},
						"owners": [
							{
								"id": 1,
								"type": "user",
								"name": "admin",
								"description": "",
								"url": "/api/v2/users/1/"
							}
						]
					},
					"created": "2021-01-11T17:37:27.381945Z",
					"modified": "2021-01-11T17:37:27.381975Z",
					"name": "Demo Credential",
					"description": "",
					"organization": null,
					"credential_type": 1,
					"managed_by_tower": false,
					"inputs": {
						"username": "admin"
					},
					"kind": "ssh",
					"cloud": false,
					"kubernetes": false
				}
			]
		}`)

	// MockedCreateCredentialResponse mocked `CreateCredential` endpoint response
	MockedCreateCredentialResponse = []byte(`
		{
			"id": 1,
			"type": "credential",
			"url": "/api/v2/credentials/1/",
			"related": {
				"created_by": "/api/v2/users/1/",
				"modified_by": "/api/v2/users/1/",
				"activity_stream": "/api/v2/credentials/1/activity_stream/",
				"access_list": "/api/v2/credentials/1/access_list/",
				"object_roles": "/api/v2/credentials/1/object_roles/",
				"owner_users": "/api/v2/credentials/1/owner_users/",
				"owner_teams": "/api/v2/credentials/1/owner_teams/",
				"copy": "/api/v2/credentials/1/copy/",
				"input_sources": "/api/v2/credentials/1/input_sources/",
				"credential_type": "/api/v2/credential_types/1/",
				"user": "/api/v2/users/1/"
			},
			"summary_fields": {
				"credential_type": {
					"id": 1,
					"name": "Machine",
					"description": ""
				},
				"created_by": {
					"id": 1,
					"username": "admin",
					"first_name": "",
					"last_name": ""
				},
				"modified_by": {
					"id": 1,
					"username": "admin",
					"first_name": "",
					"last_name": ""
				},
				"object_roles": {
					"admin_role": {
						"description": "Can manage all aspects of the credential",
						"name": "Admin",
						"id": 19
					},
					"use_role": {
						"description": "Can use the credential in a job template",
						"name": "Use",
						"id": 20
					},
					"read_role": {
						"description": "May view settings for the credential",
						"name": "Read",
						"id": 21
					}
				},
				"user_capabilities": {
					"edit": true,
					"delete": true,
					"copy": true,
					"use": true
				},
				"owners": [
					{
						"id": 1,
						"type": "user",
						"name": "admin",
						"description": "",
						"url": "/api/v2/users/1/"
					}
				]
			},
			"created": "2021-01-11T17:37:27.381945Z",
			"modified": "2021-01-11T17:37:27.381975Z",
			"name": "Demo Credential",
			"description": "",
			"organization": null,
			"credential_type": 1,
			"managed_by_tower": false,
			"inputs": {
				"username": "admin"
			},
			"kind": "ssh",
			"cloud": false,
			"kubernetes": false
		}`)

	// MockedUpdateCredentialResponse mocked `UpdateCredential` endpoint response
	MockedUpdateCredentialResponse = []byte(`
		{
			"id": 1,
			"type": "credential",
			"url": "/api/v2/credentials/1/",
			"related": {
				"created_by": "/api/v2/users/1/",
				"modified_by": "/api/v2/users/1/",
				"activity_stream": "/api/v2/credentials/1/activity_stream/",
				"access_list": "/api/v2/credentials/1/access_list/",
				"object_roles": "/api/v2/credentials/1/object_roles/",
				"owner_users": "/api/v2/credentials/1/owner_users/",
				"owner_teams": "/api/v2/credentials/1/owner_teams/",
				"copy": "/api/v2/credentials/1/copy/",
				"input_sources": "/api/v2/credentials/1/input_sources/",
				"credential_type": "/api/v2/credential_types/1/",
				"user": "/api/v2/users/1/"
			},
			"summary_fields": {
				"credential_type": {
					"id": 1,
					"name": "Machine",
					"description": ""
				},
				"created_by": {
					"id": 1,
					"username": "admin",
					"first_name": "",
					"last_name": ""
				},
				"modified_by": {
					"id": 1,
					"username": "admin",
					"first_name": "",
					"last_name": ""
				},
				"object_roles": {
					"admin_role": {
						"description": "Can manage all aspects of the credential",
						"name": "Admin",
						"id": 19
					},
					"use_role": {
						"description": "Can use the credential in a job template",
						"name": "Use",
						"id": 20
					},
					"read_role": {
						"description": "May view settings for the credential",
						"name": "Read",
						"id": 21
					}
				},
				"user_capabilities": {
					"edit": true,
					"delete": true,
					"copy": true,
					"use": true
				},
				"owners": [
					{
						"id": 1,
						"type": "user",
						"name": "admin",
						"description": "",
						"url": "/api/v2/users/1/"
					}
				]
			},
			"created": "2021-01-11T17:37:27.381945Z",
			"modified": "2021-01-11T17:37:27.381975Z",
			"name": "Demo Credential",
			"description": "Demo Credential Updated",
			"organization": null,
			"credential_type": 1,
			"managed_by_tower": false,
			"inputs": {
				"username": "admin"
			},
			"kind": "ssh",
			"cloud": false,
			"kubernetes": false
		}`)

	// MockedDeleteCredentialResponse mocked `DeleteCredential` endpoint response
	MockedDeleteCredentialResponse = []byte(`{}`)
)
