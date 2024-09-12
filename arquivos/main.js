import express from 'express'
import * as fs from 'fs/promises'
import { createReadStream } from 'fs';
import * as config from 'dotenv';
import multer from 'multer'
import { db } from './database.js';
config.config()
const FILE_CATEGORYS = {
    PERSON: 1,
    STORE: 2,
    PRODUCT: 3,
    BUSINESS: 4
}
const ENTITYS_TABLE = {
    PERSON: "person",
    STORE: "store",
    PRODUCT: "product",
    BUSINESS: "businesses"
}
const storage = multer.diskStorage({
    destination: (req, file, cb) => {
        cb(null, `/tmp/`);
    },
    filename: (req, file, cb) => {
        const {category, id } = req.query
        if(!category || !id) {
            return res.send(404)
        }
        const fileName = `${file.originalname}`
        cb(null, fileName)
    }
})
const authMiddleware = async (req, res, next) => {
    const token = req.headers.token
    const request = await fetch("http://loja-auth/verify-token", {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({token})
    })
    if(request.status != 200) {
        return res.sendStatus(401)
    }
    next()
}
const upload = multer({storage})
const main = async () => {
    const server = express()
    server.use(express.json())
    server.use("/", authMiddleware)
    server.get('/ping', (req, res) => {
        res.send("ping-pong")
    })
    server.post('/files/upload', upload.single('file'), async (req, res) => {
        try {
            const file = req.file
            const {category, id} = req.query
            const validCategorys = ["business", "product", "store", "person"]
            if(!category || !id || !validCategorys.includes(category)) {
                return res.send(400)
            }
            const fileName = `${file.originalname}`
            const file_dest = `${fileName}`
            await fs.copyFile("/tmp/" + fileName, "/app/storage/" + file_dest)
            await fs.unlink("/tmp/" + fileName)
            const insertConfig = { file_name: file_dest }
            insertConfig[`${category}_id`] = id
            const result = await db("files").insert(insertConfig).returning('id')
            const [inserted] = result
            const inserted_id = inserted.id
            await db("file_category").insert({
                file_id: inserted_id,
                category_id: FILE_CATEGORYS[category.toUpperCase()]
            })
            res.sendStatus(200)
        } catch (error) {
            if(error.code == 23505) {// unique key violation
                return res.sendStatus(400)
            }
            console.log(error)
            res.sendStatus(500)
        }
    })
    server.get('/files/:category/:entityid', async (req, res) => {
        try {
            const {category, entityid} = req.params
            const categoryId = FILE_CATEGORYS[category.toUpperCase()]
            const entityTable = ENTITYS_TABLE[category.toUpperCase()]
            const query = `
            SELECT f.file_name, f.id
            FROM files f
            INNER JOIN file_category fc ON fc.file_id = f.id AND fc.category_id = ?
                INNER JOIN ?? et ON et.id = ?;
            `;

            const files = await db.raw(query, [categoryId, entityTable, entityid]);
            return res.json({
                status: true,
                data: files.rows,
                meta: {
                    page: 1,
                    total: files.rows.length
                }
            })
        } catch (error) {
            console.log(error)
            res.sendStatus(500)
        }
    })
    server.get('/files/:category/:entityid/:fileid', async (req, res) => {
        const { category, entityid, fileid } = req.params
        const query = `SELECT f.id, f.file_name FROM files f
        INNER JOIN file_category fc on fc.category_id = ?
        INNER JOIN ?? et on et.id = ?
        WHERE f.id = ?`
        const categoryId = FILE_CATEGORYS[category.toUpperCase()]
        const entityTable = ENTITYS_TABLE[category.toUpperCase()]
        const result = await db.raw(query, [
            categoryId, entityTable, entityid, fileid
        ])
        if(result.rows.length < 1) {
            return res.sendStatus(400)
        }
        const file_path = `/app/storage/${result.rows[0].file_name}`;
        res.setHeader('Content-Type', 'application/octet-stream')
        createReadStream(file_path).pipe(res)
    })
    server.listen(process.env.API_PORTA, () => console.log("File service running on " + process.env.API_PORTA))
}
main();
