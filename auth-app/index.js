fetch(`/public/authInfo.json`)
    .then(res => res.json())
    .then(authInfo => initKeycloak(authInfo));

function initKeycloak(authInfo) {

    const keycloak = new Keycloak({
        url: window.location.origin,
        realm: authInfo.realm,
        clientId: authInfo.client
    });

    keycloak.init({
        onLoad: 'check-sso',
        silentCheckSsoRedirectUri: window.location.origin + '/silent-check-sso.html'
    })
        .then(function (authenticated) {
            if (authenticated) {
                setAuthenticated(true)
                welcomeUser(keycloak)
            } else {
                setAuthenticated(false)
            }
        }).catch(err => {
            alert(err.error);
        });

    document.getElementById("login-btn").addEventListener('click', () => keycloak.login({ redirectUri: window.location.origin }))
    document.getElementById("logout-btn").addEventListener('click', () => keycloak.logout({ redirectUri: window.location.origin }))
}

function welcomeUser(keycloak) {
    const { idTokenParsed } = keycloak;

    document.getElementById("welcome-title").innerText = `Welcome, ${idTokenParsed.preferred_username}`;
    document.getElementById("token-p").innerText = JSON.stringify(idTokenParsed, null, 4);
}

function setAuthenticated(isAuthenticated) {
    document.getElementById("guest").hidden = isAuthenticated
    document.getElementById("authenticated").hidden = !isAuthenticated
}