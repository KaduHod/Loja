import express from 'express'
import { readFile } from 'fs/promises'
const main = async () => {
    let env_file = await readFile('.env', 'utf-8')
    const env = {}
    for(const line of env_file.split('/n')) {
        const [key, value] = line.split("=")
        env[key] = value
    }
    const server = express()
    server.use(express.json())
    server.get('/ping', (req, res) => {
        res.send("ping-pong")
    })
    server.listen(env.API_PORTA, () => console.log("Auth service running on " + env.API_PORTA))
}
main();
