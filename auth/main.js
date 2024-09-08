import express from 'express'
import * as config from 'dotenv';
import jwt from 'jsonwebtoken'
config.config()
const {env} = process
const services = {
    api: {
        secret: env.API_SERVICE_SECRET
    },
    arquivos: {
        secret: env.ARQUIVOS_SERVICE_SECRET
    }
}
const checkService = ({service, secret}) => {
    const serviceIsValid = !!services[service]
    console.log({services, service, secret})
    if(!serviceIsValid) return false
    if(services[service].secret !== secret) return false;
    return true;
}
const main = async () => {
    const server = express()
    const JWT_SECRET = env.JWT_SECRET;
    server.use(express.json())
    server.get('/ping', (req, res) => {
        res.send("pong")
    })
    server.post('/create-token', (req, res) => {
        const { service, secret } = req.body;

        if (checkService({service, secret})) {
            // Cria o token JWT
            const token = jwt.sign({ service: service }, JWT_SECRET, {
                expiresIn: '1h', // Expira em 1 hora
            });

            return res.json({ token });
        } else {
            return res.status(401).json({ message: 'Invalid services!' });
        }
    })
    server.post('/verify-token', (req, res) => {
        const { token } = req.body;
        if (!token) {
            return res.status(400).json({ message: 'Token is required' });
        }

        try {
            // Verifica e decodifica o token JWT
            jwt.verify(token, JWT_SECRET);
            return res.json({ valid: true, decoded });
        } catch (err) {
            return res.status(401).json({ valid: false, message: 'Invalid or expired token' });
        }
    })
    server.listen(env.API_PORTA, () => console.log("Auth service running on " + env.API_PORTA))
}
main();
