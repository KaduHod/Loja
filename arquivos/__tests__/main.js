import { readFile } from 'fs/promises'

require('dotenv').config()
let AUTH_HEADER = {token:""}

test('Should save a file from request', async () => {
    const file = await readFile('/app/storage/test.jpeg')
    const form = new FormData()
    form.append('file', new Blob([file], {type: "image/jpeg"}), Date.now()+"teste.jpeg")
    const request = await fetch('http://localhost:4444/files/upload?category=business&id=1', {
        method: "POST", body: form, headers: {...AUTH_HEADER}
    })
    expect(request.status).toBe(200)
})
test("Should get files from entity", async () => {
    const request = await fetch("http://localhost:4444/files/business/1", { headers: {...AUTH_HEADER} })
    expect(request.status).toBe(200)
    const res = await request.json()
    expect(res.status).toBe(true)
})
test("Should download file from server", async () => {
    const request = await fetch("http://localhost:4444/files/business/1/39", { headers: {...AUTH_HEADER} })
    expect(request.status).toBe(200)
})
beforeAll(async () => {
    const request = await fetch("http://loja-auth/create-token", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            service: "arquivos",
            secret: "arquivos-service"
        })
    })
    const {token} = await request.json()
    AUTH_HEADER = {token}
})
afterAll(async () => {

})
