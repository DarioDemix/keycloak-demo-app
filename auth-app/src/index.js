const loginBtn = document.querySelector("#login-btn");

fetchAuthInfo().then(authInfo => {
    const {realm, client, clientSecret} = authInfo;

    const openIdConnectUrl = `/realms/${realm}/protocol/openid-connect`;
    const loginURL = `${openIdConnectUrl}/auth?response_type=code&client_id=${client}&redirect_uri=${window.location.origin}`;

    const queryParams =  parseQueryParams()
   
    if (queryParams) {
        fetchToken(queryParams, openIdConnectUrl, authInfo)
            .then(res => res.json())
            .then(json => {
                const accessToken = json.access_token;
                if (accessToken) {
                    console.log(`ACCESS TOKEN: ${accessToken}`);
                    const payload = parseJwt(accessToken)

                    loginBtn.remove()
                    document.getElementById("welcome-title").innerText = `Welcome, ${payload.preferred_username}`
                    
                    document.getElementById("token-p").innerText = JSON.stringify(payload, null, 4);

                }
            })
    } else {
        bindListeners(loginURL)
    }
});

function fetchToken(queryParams, openIdConnectUrl, authInfo) {
    const {client, clientSecret} = authInfo
    const {sessionState, code} = queryParams

    const tokenAPI = `${openIdConnectUrl}/token`;

    console.log(`CODE: ${code}`);

    const params = {
        "code": code,
        "client_id": client,
        "client_secret": clientSecret,
        "redirect_uri": window.location.origin,
        "grant_type": "authorization_code"
    }

    return fetch(
        tokenAPI,
        { 
            method: "POST",
            body: buildBody(params),
        }
    )
}

function parseQueryParams() {
    if (rawParams = window.location.href.split("?")[1]) {
        const queryParams = new URLSearchParams(rawParams);
        return {
            sessionState: queryParams.get("session_state"),
            code: queryParams.get("code")
        }
    }
    console.log("No params found")
}

async function fetchAuthInfo() {
    const res = await fetch(`/info`)
    return await res.json()
    
};

function bindListeners(loginURL) {
    loginBtn.addEventListener('click', () => window.location.replace(loginURL));
}

function buildBody(params) {
    const body = new URLSearchParams();
    for (const key in params) {
        body.append(key, params[key]);
    }

    return body;
}

function parseJwt(token) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace('-', '+').replace('_', '/');
    return JSON.parse(atob(base64));
  }