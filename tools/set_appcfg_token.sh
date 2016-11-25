#!/bin/bash

access_token=`gcloud auth print-access-token`

cat <<EOF > ~/.appcfg_oauth2_tokens
{
    "_module": "oauth2client.client",
    "token_expiry": null,
    "access_token": "${access_token}",
    "token_uri": "https://accounts.google.com/o/oauth2/token",
    "invalid": false,
    "token_response": null,
    "client_id": null,
    "id_token": null,
    "client_secret": null,
    "revoke_uri": null,
    "_class": "OAuth2Credentials",
    "refresh_token": null,
    "user_agent": null
}
EOF