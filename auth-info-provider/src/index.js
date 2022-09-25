const express = require('express');

const app = express();

const { REALM, CLIENT, CLIENT_SECRET } = process.env;
const PORT = 8080;

console.log(REALM, CLIENT, CLIENT_SECRET);

app.get('/info', (req, res) => {
    res.json({
        realm: REALM,
        client: CLIENT,
        clientSecret: CLIENT_SECRET,
    });
})

app.listen(PORT, () => console.log(`App listening on port ${PORT}`));