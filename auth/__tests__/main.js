require('dotenv').config()

test('Generate valid token', async () => {
    const response = await fetch('http://localhost:5555/create-token', {
        method: "POST",
        headers: {'Content-type': 'application/json'},
        body: JSON.stringify({
            service: 'api',
            secret: process.env.API_SERVICE_SECRET
        })
    });
    expect(response.status).toBe(200);
});

test('Should return invalid credentials', async () => {
    const response = await fetch('http://localhost:5555/create-token', {
        method: "POST",
        headers: {'Content-type': 'application/json'},
        body: JSON.stringify({
            service: 'api',
            secret: "sadljsadklja"
        })
    });

    expect(response.status).toBe(400);
});

test('Should validate token', async () => {
    const createTokenResponse = await fetch('http://localhost:5555/create-token', {
        method: "POST",
        headers: {'Content-type': 'application/json'},
        body: JSON.stringify({
            service: 'api',
            secret: process.env.API_SERVICE_SECRET
        })
    });

    expect(createTokenResponse.status).toBe(200);
    const { token } = await createTokenResponse.json();
    const verifyTokenResponse = await fetch('http://localhost:5555/verify-token', {
        method: "POST",
        headers: {'Content-type': 'application/json'},
        body: JSON.stringify({ token })
    });
    expect(verifyTokenResponse.status).toBe(200);
});

test('Should not validate token', async () => {
    const response = await fetch('http://localhost:5555/verify-token', {
        method: "POST",
        headers: {'Content-type': 'application/json'},
        body: JSON.stringify({ token: "sadlkjaskdlj" })
    });

    expect(response.status).toBe(401);
});
